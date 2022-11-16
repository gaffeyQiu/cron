package goroutine

import (
	"bytes"
	"log"
	"runtime"
	"strconv"
)

func Run(f func()) {
	go runSafe(f)
}

func runSafe(f func()) {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()

	f()
}

func Id() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	// if error, just return 0
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return n
}
