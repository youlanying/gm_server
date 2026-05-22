package helper

import (
	"fmt"
	"runtime"
)

func BackTrace(tag string) {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	fmt.Printf("[BackTrace %s] %s\n", tag, string(buf[:n]))
}
