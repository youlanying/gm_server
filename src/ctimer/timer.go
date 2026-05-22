package ctimer

import (
	"bytes"
	"encoding/gob"
	"gm_server/src/logger"
	"sync/atomic"
	"time"
)

const (
	MAX_TIMER_CHAN_SIZE = 65535
)

var (
	C_TIMER     chan int64
	TIMER_MAP   map[int64]*CTimer
	nextTimerId int64
)

func init() {
	TIMER_MAP = make(map[int64]*CTimer)
	C_TIMER = make(chan int64, MAX_TIMER_CHAN_SIZE)
	nextTimerId = int64(0)
}

func expire(timerId int64) {
	C_TIMER <- timerId
}

//timer调用完成后执行，不再往deleted_map里添加，因为已经没有触发点
func CompleteTimer(timerId int64) {
	delete(TIMER_MAP, timerId)
}

func FreeTimer(timerId int64) {
	if CheckTimerExist(timerId) {
		del(timerId)
		delete(TIMER_MAP, timerId)
	} else {
		logger.LogWarnf("Free not exist timer %d.", timerId)
	}
}

// 创建定时器。delay 的单位为毫秒。
func CreateTimer(userId int64, delay int64, data CTData) int64 {
	if len(C_TIMER) >= MAX_TIMER_CHAN_SIZE {
		logger.LogErrf("Timer overflow, cur size: %v, discard timer %v %v %v.",
			len(C_TIMER), userId, delay, data)
		return 0
	}

	if delay < 0 {
		delay = 0
	}

	expired := time.Now().UnixNano()/1e6 + delay
	oldId := nextTimerId
	tid := atomic.AddInt64(&nextTimerId, 1)
	if oldId >= tid {
		logger.LogErrf("RELOOP: OLD %d >= NEW %d.", oldId, tid)
	}

	index := add(tid, expired, C_TIMER)

	TIMER_MAP[tid] = &CTimer{
		Receiver: userId,
		Expired:  expired,
		Index:    index,
		CTData:   data,
	}

	return tid
}

func CancelTimer(timerId int64) {
	FreeTimer(timerId)
}

// Marshal/unmarshal timers
type DbTimer struct {
	Receiver int64
	CTData   CTData
	Expired  int64
	Id       int64
}

type TimerInfo struct {
	Timers []DbTimer
	NextId int64
}

func LoadTimers(data []byte) (error, func()) {
	info := &TimerInfo{
		Timers: make([]DbTimer, 0),
		NextId: 0,
	}
	enc := gob.NewDecoder(bytes.NewReader(data))
	if err := enc.Decode(info); err != nil {
		logger.LogFatal("Load GS_MAP error:", err.Error())
		return err, nil
	}

	// reset next timer id
	atomic.StoreInt64(&nextTimerId, info.NextId)
	logger.Logf("load timerNextId=%d\n", nextTimerId)

	return nil, func() {
		rescheduleTimers(info)
	}
}

func rescheduleTimers(info *TimerInfo) error {
	for _, item := range info.Timers {
		// reschedule
		index := add(item.Id, item.Expired, C_TIMER)
		TIMER_MAP[item.Id] = &CTimer{
			Receiver: item.Receiver,
			Index:    index,
			Expired:  item.Expired,
			CTData:   item.CTData,
		}
		logger.Logf("reschedule timer tid=%d, ownerId=%d\n", item.Id, item.Receiver)
	}
	return nil
}

func DumpTimers() (error, []byte) {
	logger.Logf("save timerNextId=%d", atomic.LoadInt64(&nextTimerId))
	info := TimerInfo{
		Timers: make([]DbTimer, 0),
		NextId: nextTimerId,
	}

	for tid, c_timer := range TIMER_MAP {
		//c_timer.Receiver == -1 叛军-1
		if c_timer.Receiver != 0 {
			entry := DbTimer{
				Receiver: c_timer.Receiver,
				CTData:   c_timer.CTData,
				Expired:  c_timer.Expired,
				Id:       tid,
			}
			info.Timers = append(info.Timers, entry)
		}
	}

	buffer := new(bytes.Buffer)
	enc := gob.NewEncoder(buffer)
	if err := enc.Encode(info); err != nil {
		return err, nil
	}
	return nil, buffer.Bytes()
}

func CheckTimerExist(timerId int64) bool {
	if timerId == 0 {
		return true
	}
	_, has := TIMER_MAP[timerId]
	if has {
		return true
	}
	return false
}
