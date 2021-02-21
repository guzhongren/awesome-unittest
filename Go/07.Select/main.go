package racer

import (
	"fmt"
	"net/http"
	"time"
)

// Racer is a http racer
func Racer(url1, url2 string) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %v and %v", url1, url2)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
