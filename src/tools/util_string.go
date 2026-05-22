package tools

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	//将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	return md5str2
}

func Tostring(inter interface{}) (str string) {
	switch inter.(type) {
	case string:
		str = inter.(string)
		break
	case int:
		str = strconv.FormatInt(int64(inter.(int)), 10)
		break
	case int32:
		str = strconv.FormatInt(int64(inter.(int32)), 10)
		break
	case int64:
		str = strconv.FormatInt(int64(inter.(int64)), 10)
		break
	case uint32:
		str = strconv.FormatUint(uint64(inter.(uint32)), 10)
		break
	case uint64:
		str = strconv.FormatUint(inter.(uint64), 10)
		break
	}
	return str
}
