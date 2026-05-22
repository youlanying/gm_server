package network

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	"gm_server/src/bridge"
	"gm_server/src/logger"
	"gm_server/src/tools/packet"
	"net"
)

const (
	DEFAULT_QUEUE_SIZE      = 256
	DEFAULT_CLOSE_CHAN_SIZE = 8
)

type preMsg struct {
	cmd  string
	body []byte
}

type Sender struct {
	exitChannel chan bool
	msgQueue    chan []byte
	maxSize     int
	connect     *net.TCPConn
}

func newSender(conn *net.TCPConn, pkgChanSize int32) *Sender {
	size := DEFAULT_QUEUE_SIZE
	if pkgChanSize > 0 {
		size = int(pkgChanSize)
	}
	return &Sender{
		connect:     conn,
		exitChannel: make(chan bool, DEFAULT_CLOSE_CHAN_SIZE),
		maxSize:     size,
		msgQueue:    make(chan []byte, size),
	}
}

func (self *Sender) start() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Sender %v start Error : %v", self, err)
		}
	}()
	for {
		select {
		case data := <-self.msgQueue:
			self.rawSend(data)
		case <-self.exitChannel:
			fmt.Printf("Sender %p close !: %v \n", self, self)
			return
		}
	}
}

func (self *Sender) send(pbData proto.Message) (err error) {
	if self == nil {
		logger.LogErr("sender is nil, msg:", pbData)
		return nil
	}
	if len(self.msgQueue) < self.maxSize {
		data, err := bridge.GetFSMessageData(pbData)
		if err != nil {
			logger.LogErrf("send msg marshaling error: %v , pbData = [%v]", err, pbData)
			return err
		}
		self.msgQueue <- data
		return nil
	} else {
		logger.LogWarnf("Sender overflow, msgQueue: %d addr: %v", len(self.msgQueue), self.connect.RemoteAddr())
		return errors.New(fmt.Sprintf("Sender overflow, addr: %v", self.connect.RemoteAddr()))
	}
	return nil
}

func (self *Sender) syncSend(pbData proto.Message) (err error) {
	data, err := bridge.GetFSMessageData(pbData)
	if err != nil {
		logger.Logf("send msg marshaling error: %v", err)
		return err
	}
	return self.rawSend(data)
}

func (self *Sender) sendBytes(data []byte) (err error) {
	if len(self.msgQueue) < self.maxSize {
		self.msgQueue <- data
		return nil
	} else {
		logger.LogWarnf("Sender overflow, msgQueue: %d addr: %v", len(self.msgQueue), self.connect.RemoteAddr())
		return errors.New(fmt.Sprintf("Sender overflow, addr: %v", self.connect.RemoteAddr()))
	}
	return nil
}

func (self *Sender) syncSendBytes(data []byte) (err error) {
	return self.rawSend(data)
}

func (self *Sender) rawSend(data []byte) error {
	write := packet.Writer()
	dataSize := len(data)
	if dataSize > MAX_SIZE {
		logger.LogErrf("Sender overflow, size: %d.", dataSize)
		return errors.New(fmt.Sprintf("Sender overflow, size: %v", dataSize))
	}
	write.WriteU32(uint32(dataSize))
	write.WriteRawBytes(data)
	_, err := self.connect.Write(write.Data())
	if err != nil {
		logger.LogWarnf("Error send replay: %v, sender %v ,bufsize %v msgQueue size:%v",
			err, self, dataSize, len(self.msgQueue))

		return err
	}
	return nil
}

func (self *Sender) close() {
	self.exitChannel <- true
}

////////////////////////////////////////kcp部分//////////////////////////////
type KcpSender struct {
	exitChannel chan bool
	msgQueue    chan []byte
	maxSize     int
	kcpConnect  *kcp.UDPSession
}

func newKcpSender(conn *kcp.UDPSession, pkgChanSize int32) *KcpSender {
	size := DEFAULT_QUEUE_SIZE
	if pkgChanSize > 0 {
		size = int(pkgChanSize)
	}
	return &KcpSender{
		kcpConnect:  conn,
		exitChannel: make(chan bool, DEFAULT_CLOSE_CHAN_SIZE),
		maxSize:     size,
		msgQueue:    make(chan []byte, size),
	}
}

func (self *KcpSender) start() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Sender %v start Error : %v", self, err)
		}
	}()
	for {
		select {
		case data := <-self.msgQueue:
			self.rawSend(data)
		case <-self.exitChannel:
			fmt.Printf("Sender %p close !: %v \n", self, self)
			return
		}
	}
}

func (self *KcpSender) rawSend(data []byte) error {
	write := packet.Writer()
	dataSize := len(data)
	if dataSize > MAX_SIZE {
		logger.LogErrf("Sender overflow, size: %d.", dataSize)
		return errors.New(fmt.Sprintf("Sender overflow, size: %v", dataSize))
	}
	write.WriteU32(uint32(dataSize))
	write.WriteRawBytes(data)
	_, err := self.kcpConnect.Write(write.Data())
	if err != nil {
		logger.LogWarnf("Error send replay: %v, sender %v ,bufsize %v msgQueue size:%v",
			err, self, dataSize, len(self.msgQueue))

		return err
	}
	return nil
}

func (self *KcpSender) close() {
	self.exitChannel <- true
}

func (self *KcpSender) send(pbData proto.Message) (err error) {
	if len(self.msgQueue) < self.maxSize {
		data, err := bridge.GetFSMessageData(pbData)
		if err != nil {
			logger.LogErrf("send msg marshaling error: %v , pbData = [%v]", err, pbData)
			return err
		}
		self.msgQueue <- data
		return nil
	} else {
		logger.LogWarnf("Sender overflow, msgQueue: %d addr: %v", len(self.msgQueue), self.kcpConnect.RemoteAddr())
		return errors.New(fmt.Sprintf("Sender overflow, addr: %v", self.kcpConnect.RemoteAddr()))
	}
	return nil
}

func (self *KcpSender) syncSend(pbData proto.Message) (err error) {
	data, err := bridge.GetFSMessageData(pbData)
	if err != nil {
		logger.Logf("KcpSender send msg marshaling error: %v", err)
		return err
	}
	return self.rawSend(data)
}

func (self *KcpSender) sendBytes(data []byte) (err error) {
	if len(self.msgQueue) < self.maxSize {
		self.msgQueue <- data
		return nil
	} else {
		logger.LogWarnf("Sender overflow, msgQueue: %d addr: %v", len(self.msgQueue), self.kcpConnect.RemoteAddr())
		return errors.New(fmt.Sprintf("Sender overflow, addr: %v", self.kcpConnect.RemoteAddr()))
	}
	return nil
}
