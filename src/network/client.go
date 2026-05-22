package network

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	"gm_server/src/logger"
	"net"
)

type ConnectInfo struct {
	connection     *net.TCPConn
	sender         *Sender
	kickout        bool
	ExitChannel    chan int
	MessageChannel chan *NetMessage
}

func ConnectServer(addr string, psize int32) (error, *ConnectInfo) {
	logger.Logf("Conn server %v psize %v.", addr, psize)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		return err, nil
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err, nil
	}

	info := ConnectInfo{
		connection: conn,
		sender:     newSender(conn, psize),
		kickout:    false,

		ExitChannel:    make(chan int, 32),
		MessageChannel: make(chan *NetMessage, MSG_LIST_SIZE),
	}

	go info.start()

	return nil, &info
}

func (self *ConnectInfo) Stop() {
	self.kickout = true
	self.ExitChannel <- 1
	if self.sender != nil {
		self.sender.close()
	}
	if self.connection != nil {
		self.connection.Close()
	}
}

func (self *ConnectInfo) Send(pbData proto.Message) error {
	return self.sender.send(pbData)
}

func (self *ConnectInfo) SyncSend(pbData proto.Message) error {
	return self.sender.syncSend(pbData)
}

func (self *ConnectInfo) start() {
	defer func() {
		logger.Logf("Connection to server %p %v quit.", self, self)
		self.connection.Close()
		self.sender.close()
		self.ExitChannel <- 1
	}()

	conn := self.connection
	logger.Logf("Connection %p start.", self)

	go self.sender.start()

	header := make([]byte, 4)
	for {
		if self.kickout {
			logger.LogWarnf("Kickout, conn %p %v quit.", self, self)
			return
		}

		data := readMassage(conn, header)
		if data == nil {
			logger.LogWarnf("Read msg fail, conn %p %v quit.", self, self)
			return
		}

		msg := parseMassage(data, nil) // 作为客户端连接的消息，Session为NIL。
		if msg == nil {
			logger.LogWarnf("Parse msg fail, conn %p %v quit.", self, self)
			return
		}
		self.MessageChannel <- msg
	}
}

////////////////////kcp部分///////////////////////////////////////
type KcpConnectInfo struct {
	kcpConnection  *kcp.UDPSession
	kcpSender      *KcpSender
	kickout        bool
	ExitChannel    chan int
	MessageChannel chan *NetMessage
}

func ConnectKcpServer(addr string, psize int32) (error, *KcpConnectInfo) {
	logger.Logf("Conn kcp server %v psize %v.", addr, psize)
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		return err, nil
	}

	conn, err := kcp.DialWithOptions(udpAddr.String(), nil, 10, 3)
	//conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err, nil
	}

	info := KcpConnectInfo{
		kcpConnection:  conn,
		kcpSender:      newKcpSender(conn, psize),
		kickout:        false,
		ExitChannel:    make(chan int, 32),
		MessageChannel: make(chan *NetMessage, MSG_LIST_SIZE),
	}

	go info.start()

	return nil, &info
}

func (self *KcpConnectInfo) start() {
	defer func() {
		logger.Logf("Connection to kcp server %p %v quit.", self, self)
		self.kcpConnection.Close()
		self.kcpSender.close()
		self.ExitChannel <- 1
	}()

	conn := self.kcpConnection
	logger.Logf("Connection %p start.", self)

	go self.kcpSender.start()

	header := make([]byte, 4)
	for {
		if self.kickout {
			logger.LogWarnf("Kickout, conn %p %v quit.", self, self)
			return
		}

		data := readKcpMassage(conn, header)
		if data == nil {
			logger.LogWarnf("Read msg fail, conn %p %v quit.", self, self)
			return
		}

		msg := parseMassage(data, nil) // 作为客户端连接的消息，Session为NIL。
		if msg == nil {
			logger.LogWarnf("Parse msg fail, conn %p %v quit.", self, self)
			return
		}
		self.MessageChannel <- msg
	}
}

func (self *KcpConnectInfo) Send(pbData proto.Message) error {
	return self.kcpSender.send(pbData)
}
