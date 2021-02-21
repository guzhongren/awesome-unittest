package mocking

import (
	"fmt"
	"io"
	"time"
)

// Countdown 向下计数
func Countdown(w io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(w, "go")
}
