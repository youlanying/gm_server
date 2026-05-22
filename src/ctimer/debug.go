package ctimer

import (
	"gm_server/src/helper"
	"gm_server/src/logger"
)

type timerState struct {
	id   int64
	site int
}

const (
	_in_queue = 1000000
)

func GetTriggerMap() map[int64]*timerState {
	triggerMap := make(map[int64]*timerState)
	for _, event := range _eventqueue {
		triggerMap[event.Id] = &timerState{
			id:   event.Id,
			site: _in_queue,
		}
	}
	for i := int(TIMER_LEVEL - 1); i >= 0; i-- {
		list := _eventlist[i]
		for _, event := range list {
			triggerMap[event.Id] = &timerState{
				id:   event.Id,
				site: i,
			}
		}
	}
	return triggerMap
}

func PrintNoTriggerTimer() {
	defer func() {
		if err := recover(); err != nil {
			helper.BackTrace("worker")
			logger.LogErrf("%+v.", err)
		}
	}()
	triggerMap := GetTriggerMap()

	logger.Logf("Event count:%d.", len(triggerMap))
	for timerId, t := range TIMER_MAP {
		if _, exist := triggerMap[timerId]; !exist {
			logger.LogErrf("Timer:%+v no trigger.", t)
		}
	}
}

func PrintNoTriggerDeletedTimer() {
	triggerMap := GetTriggerMap()

	logger.LogErrf("Event count:%d,deleted %d..", len(triggerMap), len(_deleted_timers))
	for timerId, v := range _deleted_timers {
		if _, exist := triggerMap[timerId]; !exist {
			logger.LogErrf("Timer:%d %v no trigger.", timerId, v)
		}
	}
}

func RebuildNoTriggerTimer() {
	triggerMap := GetTriggerMap()

	logger.Logf("Event count:%d.", len(triggerMap))
	for timerId, t := range TIMER_MAP {
		if _, exist := triggerMap[timerId]; !exist {
			logger.LogErrf("Rebuild Timer:%+v  trigger.", t)
			add(timerId, t.Expired, C_TIMER)
		}
	}
}

func PrintTrigger(timerId int64) {
	for _, event := range _eventqueue {
		if event.Id == timerId {
			logger.Logf("Timer %d in queue event %+v.", timerId, event)
			return
		}
	}
	for i := int(TIMER_LEVEL - 1); i >= 0; i-- {
		for _, event := range _eventlist[i] {
			if event.Id == timerId {
				logger.Logf("Timer %d in list %d event %+v.", timerId, i, event)
				return
			}
		}
	}
}

func GetTimerSendCount() int64 {
	return _timer_send_cnt
}

func GetAddTimerCount() int64 {
	return _timer_id
}

func GetTriggerTimerCount() int64 {
	tm := GetTriggerMap()
	return int64(len(tm))
}

func GetChannelTimerCount() int64 {
	return int64(len(C_TIMER))
}

func GetSimpleTriggerTimerCount() int64 {
	return _trigger_cnt + int64(len(_eventqueue))
}

func GetDeletedTimerCount() int64 {
	return int64(len(_deleted_timers))
}
