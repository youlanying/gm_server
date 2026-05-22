package network

import (
	"fmt"
	"github.com/xtaci/kcp-go"
	"gm_server/src/logger"
	"net"
	"strings"
	"sync/atomic"
)

var (
	_curSessionId = int32(0)
)

const (
	TCP_TIME_OUT  = 300
	MSG_LIST_SIZE = 1<<14 - 1
)

type NetMessage struct {
	NetSession *Session
	Head       string
	Body       []byte
}

type ServerChannel struct {
	ConnectChannel chan *Session
	ExitChannel    chan *Session
	MessageChannel chan *NetMessage
}

func NewServer(address string, msgSize int32) (*ServerChannel, error) {
	serverChan := ServerChannel{
		ConnectChannel: make(chan *Session, 5),
		ExitChannel:    make(chan *Session, 5),
		MessageChannel: make(chan *NetMessage, MSG_LIST_SIZE),
	}

	tcp4Address, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		logger.LogErrf("[net] error tcp4Address:" + err.Error())
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", tcp4Address)
	if err != nil {
		logger.LogErrf("[net] error listening:" + err.Error())
		return nil, err
	}
	go func() {
		for {
			connectTcp, err := listener.AcceptTCP()
			if err != nil {
				logger.LogErrf("[net] AcceptTCP error:" + err.Error())
				continue
			}
			tempNewSession := newSession(connectTcp, &serverChan, msgSize)
			serverChan.ConnectChannel <- tempNewSession
		}
	}()

	return &serverChan, nil
}

func newSession(connectTcp *net.TCPConn, serverChan *ServerChannel, pkgChanSize int32) *Session {
	ip := net.ParseIP(strings.Split(connectTcp.RemoteAddr().String(), ":")[0])

	id := atomic.AddInt32(&_curSessionId, 1)
	tempSession := &Session{
		Id:             id,
		Ip:             ip,
		connection:     connectTcp,
		kickout:        false,
		sender:         newSender(connectTcp, pkgChanSize),
		exitChannel:    serverChan.ExitChannel,
		messageChannel: serverChan.MessageChannel,
		kcpConnection:  nil,
		kcpSender:      nil,
	}

	return tempSession
}

func StartSession(netSession *Session) {
	go netSession.start()
	go netSession.sender.start()
}

func NewTcpAndKcpServer(tcpAaddress string, kcpAaddress string, msgSize int32) (*ServerChannel, error) {
	serverChan := ServerChannel{
		ConnectChannel: make(chan *Session, 5),
		ExitChannel:    make(chan *Session, 5),
		MessageChannel: make(chan *NetMessage, MSG_LIST_SIZE),
	}
	err := NewTcpServer(tcpAaddress, msgSize, &serverChan)
	if err != nil {
		fmt.Printf("NewTcpServer fail. err = %v \n", err)
	}
	err = NewKcpServer(kcpAaddress, msgSize, &serverChan)
	if err != nil {
		fmt.Printf("NewKcpServer fail. err = %v \n", err)
	}

	return &serverChan, nil
}

func NewTcpServer(address string, msgSize int32, serverChan *ServerChannel) error {

	tcp4Address, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		logger.LogErrf("[net] error tcp4Address:" + err.Error())
		return err
	}

	listener, err := net.ListenTCP("tcp", tcp4Address)
	if err != nil {
		logger.LogErrf("[net] error listening:" + err.Error())
		return err
	}
	go func() {
		for {
			connectTcp, err := listener.AcceptTCP()
			if err != nil {
				logger.LogErrf("[net] AcceptTCP error:" + err.Error())
				continue
			}
			tempNewSession := newSession(connectTcp, serverChan, msgSize)
			serverChan.ConnectChannel <- tempNewSession
		}
	}()

	return nil
}

func NewKcpServer(address string, msgSize int32, serverChan *ServerChannel) error {

	udp4Address, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		logger.LogErrf("[net] error tcp4Address:" + err.Error())
		return err
	}

	listener, err := kcp.ListenWithOptions(udp4Address.String(), nil, 10, 3)
	//listener, err := net.ListenTCP("tcp", tcp4Address)
	if err != nil {
		logger.LogErrf("[net] error ListenWithOptions:" + err.Error())
		return err
	}
	go func() {
		for {
			connectKcp, err := listener.AcceptKCP()
			if err != nil {
				logger.LogErrf("[net] AcceptKCP error:" + err.Error())
				continue
			}
			tempNewSession := newKcpSession(connectKcp, serverChan, msgSize)
			serverChan.ConnectChannel <- tempNewSession
		}
	}()

	return nil
}

func newKcpSession(connectKcp *kcp.UDPSession, serverChan *ServerChannel, pkgChanSize int32) *Session {
	ip := net.ParseIP(strings.Split(connectKcp.RemoteAddr().String(), ":")[0])
	id := atomic.AddInt32(&_curSessionId, 1)
	tempSession := &Session{
		Id:             id,
		Ip:             ip,
		connection:     nil,
		kickout:        false,
		sender:         nil,
		exitChannel:    serverChan.ExitChannel,
		messageChannel: serverChan.MessageChannel,
		kcpConnection:  connectKcp,
		kcpSender:      newKcpSender(connectKcp, pkgChanSize),
	}

	return tempSession
}

func StartKcpSession(netSession *Session) {
	go netSession.kcpStart()
	go netSession.kcpSender.start()
}
