package netmsg

import (
	"github.com/astaxie/beego"
	"sync"
)

type Handler struct {
	Pipe chan interface{}
}

type Session struct {
	Id     int64
	Hander *Handler
}

var (
	idPool   int64
	sessions map[int64]*Session
	m        *sync.RWMutex
)

func InitSession() {
	sessions = make(map[int64]*Session)
	m = new(sync.RWMutex)
}

func (s *Session) Write(data interface{}) {
	s.Hander.Pipe <- data
}

func (s *Session) Read() interface{} {
	if msg, err := <-s.Hander.Pipe; err {
		return msg
	}
	beego.Error("Net close chan error:")
	return nil
}

func NewSession() int64 {
	m.Lock()
	defer m.Unlock()
	h := Handler{
		Pipe: make(chan interface{}),
	}
	idPool++
	s := Session{
		Id:     idPool,
		Hander: &h,
	}
	sessions[s.Id] = &s
	return s.Id
}

func GetSession(id int64) *Session {
	m.RLock()
	defer m.RUnlock()
	if s, ok := sessions[id]; ok {
		return s
	}
	return nil
}

func DelSession(id int64) {
	m.Lock()
	defer m.Unlock()
	delete(sessions, id)
}
