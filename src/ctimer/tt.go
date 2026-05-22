package ctimer

import (
	"fmt"
	"gm_server/src/logger"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type _timer_event struct {
	Id      int64      // 用户自定义id
	Timeout int64      // 到期时间 Unix Time (ms)
	CH      chan int64 // 发送通道
}

const (
	TIMER_LEVEL = uint(48) // 时间段最大分级，最大时间段为 2^TIMERLEVEL
	TICK_TIME   = 10       // 10 ms
)

var (
	_eventlist [TIMER_LEVEL]map[int64]_timer_event // 事件列表

	_eventqueue      map[int64]_timer_event // 事件添加队列
	_eventqueue_lock sync.Mutex

	_deleted_timers      map[int64]bool
	_deleted_timers_lock sync.Mutex

	_timer_id int64 // 内部事件编号

	_recover_log_times int
	_timer_send_cnt    int64 // 往GS_TIMER channel输入的timer数
	_trigger_cnt       int64 // trigger中的timer数
)

func init() {
	for k := range _eventlist {
		_eventlist[k] = make(map[int64]_timer_event)
	}

	_eventqueue = make(map[int64]_timer_event)
	_deleted_timers = make(map[int64]bool)

	_recover_log_times = 10

	go _timer()
}

func current() int64 {
	return time.Now().UnixNano() / 1e6
}

//var tmp = 1

//------------------------------------------------
// 定时器 goroutine
// 根据程序启动后经过的秒数计数
func _timer() {
	timer_count := uint64(0)
	last := current()

	for {
		time.Sleep(TICK_TIME * time.Millisecond)

		// 处理排队
		// 最小的时间间隔，处理为1ms
		_eventqueue_lock.Lock()
		usedEvents := make(map[int64]_timer_event, len(_eventqueue))
		for k, v := range _eventqueue {
			// 处理微小间隔
			diff := v.Timeout - current()
			if diff <= 0 {
				diff = 1
			}

			// 发到合适的框
			for i := TIMER_LEVEL - 1; i >= 0; i-- {
				if diff >= 1<<i {
					_eventlist[i][k] = v
					usedEvents[k] = v
					_trigger_cnt++
					break
				}
			}
		}
		if len(_eventqueue) > 0 {
			if len(_eventqueue) != len(usedEvents) {
				for k, v := range usedEvents {
					if _, exist := _eventqueue[k]; !exist {
						logger.LogAlertf("TIMER %d %+v no in _eventqueue.", k, v)
					}
				}
			}
			_eventqueue = make(map[int64]_timer_event)
		}
		_eventqueue_lock.Unlock()

		// 检查事件触发
		// 累计距离上一次触发的秒数,并逐秒触发
		// 如果校正了系统时间，时间前移，nsec为负数的时候，last的值不应该变动，否则会出现秒数的重复计数
		now := current()
		nmsecs := now - last

		if nmsecs < 0 {
			logger.LogErrf("Gstimer nmsecs now %d < last %d.", now, last) // 如果系统时间回退，会打印
			continue
		} else if nmsecs == 0 {
			continue
		} else {
			last = now
		}

		for c := int64(0); c < nmsecs; c++ {
			timer_count++

			for i := TIMER_LEVEL - 1; i > 0; i-- {
				mask := (uint64(1) << i) - 1
				if timer_count&mask == 0 {
					_trigger(i)
				}
			}

			_trigger(0)
		}
	}
}

//------------------------------------------------ 单级触发
func _trigger(level uint) {
	now := current()
	list := _eventlist[level]

	for k, v := range list {
		if v.Timeout-now < 1<<level {
			// 移动到前一个更短间距的LIST
			if level == 0 {
				func() {
					defer func() {
						if err := recover(); err != nil { // ignore closed channel
							if _recover_log_times > 0 {
								_recover_log_times--
								str := fmt.Sprintf("TRIGGER err:%+v channel len:%d.", err, len(v.CH))
								logger.LogErrf(str)
								fmt.Fprintf(os.Stderr, str)
							}
						}
					}()
					if checkValidTimer(v.Id) {
						_timer_send_cnt++
						v.CH <- v.Id
					}
					_trigger_cnt-- // 不管是否可用，都要从trigger删掉
				}()
			} else {
				_eventlist[level-1][k] = v
			}
			delete(list, k)
		}
	}
}

//------------------------------------------------
// 添加一个定时，timeout为到期的Unix时间
// id 是调用者定义的编号, 事件发生时，会把id发送到ch
func add(id int64, timeout int64, ch chan int64) int64 {

	timer_id := atomic.AddInt64(&_timer_id, 1)
	event := _timer_event{Id: id, CH: ch, Timeout: timeout}

	_eventqueue_lock.Lock()
	_eventqueue[timer_id] = event
	_eventqueue_lock.Unlock()

	return timer_id
}

func del(id int64) {
	_deleted_timers_lock.Lock()
	_deleted_timers[id] = true
	_deleted_timers_lock.Unlock()
}

func checkValidTimer(id int64) bool {
	flag := true
	_deleted_timers_lock.Lock()
	if deleted, _ := _deleted_timers[id]; deleted {
		flag = false
	}
	delete(_deleted_timers, id)
	_deleted_timers_lock.Unlock()
	return flag
}

func SetRecoverLogTimes(i int) {
	logger.Logf("Set recover log times from %d to %d.", _recover_log_times, i)
	_recover_log_times = i
}
