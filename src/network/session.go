package network

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	"gm_server/src/logger"
	"gm_server/src/network/message"
	"io"
	"net"
	"time"
)

const MAX_SIZE = 1024 * 1024 * 32

type Session struct {
	Id             int32
	Ip             net.IP
	connection     *net.TCPConn
	sender         *Sender
	kickout        bool
	exitChannel    chan *Session
	messageChannel chan *NetMessage
	kcpConnection  *kcp.UDPSession
	kcpSender      *KcpSender
}

func (self *Session) start() {
	defer func() {
		if r := recover(); r != nil {
			logger.LogErrf("Session start error: %v.", r)
		}

		logger.Logf("Session %d %p quit  kick : %v", self.Id, self, self.kickout)
		if self.kickout == false {
			self.connection.Close()
			self.sender.close()
		}
		self.kickout = true
		self.exitChannel <- self
	}()
	connect := self.connection
	header := make([]byte, 4)
	for {
		if self.kickout {
			return
		}

		data := readMassage(connect, header)

		if data == nil {
			return
		}

		tempMsg := parseMassage(data, self)
		if tempMsg == nil {
			return
		}
		self.messageChannel <- tempMsg
	}
}

func (self *Session) Send(pbData proto.Message) error {
	if !self.IsTcpConnect() {
		return self.kcpSender.send(pbData)
	}
	return self.sender.send(pbData)
}

func (self *Session) SyncSend(pbData proto.Message) error {
	if !self.IsTcpConnect() {
		return self.kcpSender.syncSend(pbData)
	}
	return self.sender.syncSend(pbData)
}

/**
FSMessage fsData
*/
func (self *Session) SendBytes(fsData []byte) error { ////////
	if !self.IsTcpConnect() {
		return self.kcpSender.sendBytes(fsData)
	}
	return self.sender.sendBytes(fsData)
}

/**
FSMessage fsData
*/
func (self *Session) SyncSendBytes(fsData []byte) error {
	return self.sender.syncSendBytes(fsData)
}

func (self *Session) Close() {
	if !self.kickout {
		if self.connection != nil {
			self.connection.Close()
			self.sender.close()
		}

		if self.kcpConnection != nil {
			self.kcpConnection.Close()
			self.kcpSender.close()
		}

		self.kickout = true
	}
}

func (self *Session) GetIpAddress() string {
	return self.Ip.String()
}

func readMassage(connect *net.TCPConn, header []byte) *[]byte {
	//connect.SetReadDeadline(time.Now().Add(TCP_TIME_OUT * time.Second)) //300秒无消息返回错误，可以使连接断开
	_, err := io.ReadFull(connect, header)
	if err != nil {
		if err.Error() == "EOF" {
			logger.LogErrf("io.ReadFull Error :EOF address = %v", connect.RemoteAddr())
		} else {
			logger.LogWarnf("io.ReadFull Error  address = %v, err = %v", connect.RemoteAddr(), err)
		}
		return nil
	}
	size := binary.BigEndian.Uint32(header)
	if size > MAX_SIZE {
		logger.LogErrf("recieve msg size > MAX_SIZE address = %v", connect.RemoteAddr())
		return nil
	}

	databuf := make([]byte, size)
	_, err = io.ReadFull(connect, databuf)
	if err != nil {
		logger.LogWarnf("io.ReadFull Err receive  data %p %v", connect, err)
		return nil
	}
	return &databuf
}

func parseMassage(data *[]byte, netSession *Session) *NetMessage {
	defer func() {
		if err := recover(); err != nil {
			logger.LogWarnf("Massage parse Error  %v %v", data, err)
		}
	}()

	fsmsg := &network_message.FSMessage{}
	err := proto.Unmarshal(*data, fsmsg)
	if err != nil {
		logger.LogWarnf("proto parse Error  %v %v", data, err)
		return nil
	}

	//解析proto
	return &NetMessage{
		NetSession: netSession,
		Head:       fsmsg.GetHead(),
		Body:       fsmsg.GetBody(),
	}
}

func (self *Session) IsTcpConnect() bool {
	if self.connection != nil {
		return true
	}
	return false
}

func (self *Session) kcpStart() {
	defer func() {
		if r := recover(); r != nil {
			logger.LogErrf("Session kcpStart error: %v.", r)
		}

		logger.Logf("Session %d %p quit  kick : %v", self.Id, self, self.kickout)
		if self.kickout == false {
			//self.connection.Close()
			//self.sender.close()
			self.kcpConnection.Close()
			self.kcpSender.close()
		}
		self.kickout = true
		self.exitChannel <- self
	}()
	kcpConnect := self.kcpConnection
	header := make([]byte, 4)
	for {
		if self.kickout {
			return
		}

		data := readKcpMassage(kcpConnect, header)

		if data == nil {
			return
		}

		tempMsg := parseMassage(data, self)
		if tempMsg == nil {
			return
		}
		self.messageChannel <- tempMsg
	}
}

func readKcpMassage(connect *kcp.UDPSession, header []byte) *[]byte {
	connect.SetReadDeadline(time.Now().Add(TCP_TIME_OUT * time.Second))
	_, err := io.ReadFull(connect, header)
	if err != nil {
		if err.Error() == "EOF" {
			logger.Logf("io.ReadFull Error :EOF %p", connect)
		} else {
			logger.LogWarnf("io.ReadFull Error  %v %v", connect, err)
		}
		return nil
	}
	size := binary.BigEndian.Uint32(header)
	if size > MAX_SIZE {
		logger.LogErrf("io.ReadFull is too HUGE..  %v", size)
		return nil
	}
	databuf := make([]byte, size)
	_, err = io.ReadFull(connect, databuf)
	if err != nil {
		logger.LogWarnf("io.ReadFull Err receive  data %p %v", connect, err)
		return nil
	}
	return &databuf
}

//func readMassage(connect *net.TCPConn, header []byte) *[]byte {
//	connect.SetReadDeadline(time.Now().Add(TCP_TIME_OUT * time.Second))
//	_, err := io.ReadFull(connect, header)
//	fmt.Printf("header = %v", header)
//	if err != nil {
//		if err.Error() == "EOF" {
//			fmt.Printf("io.ReadFull Error :EOF %p \n", connect)
//		} else {
//			fmt.Printf("io.ReadFull Error  %v %v \n", connect, err)
//		}
//		return nil
//	}
//	size := uint32(binary.LittleEndian.Uint16(header))
//	fmt.Printf("size = %v \n", size)
//	if size == 0xFFFF {
//		nHeader := make([]byte, 4)
//		fmt.Printf("io.ReadFull has HUGE MSG  %p\n", connect)
//		_, err := io.ReadFull(connect, nHeader)
//		if err != nil {
//			if err.Error() == "EOF" {
//				fmt.Printf("0xFFFF io.ReadFull Error :EOF %p \n", connect)
//			} else {
//				fmt.Printf("0xFFFF io.ReadFull Error  %v %v \n", connect, err)
//			}
//			return nil
//		}
//		size = binary.LittleEndian.Uint32(nHeader)
//		fmt.Printf("io.ReadFull has HUGE MSG size =  %v\n", size)
//		if size > MAX_SIZE {
//			fmt.Printf("io.ReadFull is too HUGE..  %v\n", size)
//			return nil
//		}
//	}
//	databuf := make([]byte, size)
//	_, err = io.ReadFull(connect, databuf)
//	if err != nil {
//		fmt.Printf("io.ReadFull Err receive  data %p %v \n", connect, err)
//		return nil
//	}
//	fmt.Printf("io.ReadFull receive  datasize %v \n", size)
//	return &databuf
//}
