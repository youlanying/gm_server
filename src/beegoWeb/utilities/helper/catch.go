// Package helper : catch.go implements boilerplate code for the web service.
package helper

import (
	"fmt"
	"runtime"
)

// CatchPanic is used to catch any Panic and log exceptions to Stdout. It will also write the stack trace
func CatchPanic(err *error, sessionID string, functionName string) {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}
}

func BackTrace(tag string) {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	fmt.Printf("[BackTrace %s] %s\n", tag, string(buf[:n]))
}
