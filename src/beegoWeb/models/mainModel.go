package models

import (
	"fmt"
	"sync"
)

/**
 * 初始化
 */
func Initialize() {

}

//---- config 词典   ----

var ConfigDic map[string]string = make(map[string]string)

func GetConfig(key string) string {

	return ConfigDic[key]
}

const (
	Special1 = "白"
	Special2 = "黑"
)

var SpecialList = [2]string{Special1, Special2}

var DisableList = [4]int{0, 1, 2, 4}

func ListsMember(tar interface{}, arr []interface{}) bool {
	for _, mem := range arr {
		if mem == tar {
			return true
		}
	}

	return false
}

func UnShift(tar interface{}, arr []interface{}) []interface{} {

	var newArr []interface{}
	newArr = append(newArr, tar)

	for _, mem := range arr {
		newArr = append(newArr, mem)
	}

	return newArr
}

func ListConant(arr []interface{}, brr []interface{}) []interface{} {

	for _, mem := range brr {
		arr = append(arr, mem)
	}

	return arr
}

func ListsSplit(N int, arr []interface{}) ([]interface{}, []interface{}) {

	var newArr []interface{}

	arrlen := len(arr)

	if N >= arrlen {
		return arr, newArr
	}

	var brr []interface{}

	for _, mem := range arr {

		if N < arrlen {
			newArr = append(newArr, mem)
		} else {
			brr = append(brr, mem)
		}
	}

	return newArr, brr
}

//---- etc etc主要是做操作锁定用的 ----

var Etc sync.Map //同步锁

//-- ets名字 --
var LOCK_ETS string = "LOCK_ETS"

var LOCK_NUMBER_CREATE string = "LOCK_NUMBER_CREATE"

func EtcInsert(etcName string, key string, val string) {

	result, ok := Etc.Load(etcName)

	if ok {
		vMap := result.(map[string]string)
		vMap[key] = val
	} else {
		newMap := make(map[string]string)
		newMap[key] = val
		Etc.Store(etcName, newMap)
	}
}

func EtcLookUp(etcName string, key string) bool {

	fmt.Println("EtcLookUp", etcName, key)

	result, ok := Etc.Load(etcName)

	if ok {
		vMap := result.(map[string]string)
		fmt.Println("EtcLookUp return ", vMap)
		return vMap[key] != ""
	}

	fmt.Println("EtcLookUp return 无")
	return false
}
