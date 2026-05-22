package stack_helper

import (
	"fmt"
	"runtime"
)

func NetWorkBackTrace(tag string, msg string) {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	fmt.Printf("[NetWorkBackTrace %s] msg: %s\nstack: %s\n", tag, msg, string(buf[:n]))
}
