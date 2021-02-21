package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/**
编写一个叫做 WebsiteRacer 的函数，
用来对比请求两个 URL 来「比赛」，
并返回先响应的 URL。
如果两个 URL 在 10 秒内都未返回结果，那么应该返回一个 error
*/

func TestRacer(t *testing.T) {
	t.Run(`should return 42`, func(t *testing.T) {
		expected := 42
		got := 42
		if expected != got {
			t.Fatal(`测试没正确执行`)
		}
	})
	t.Run("should return faster", func(t *testing.T) {
		slowServer := deleyedServer(20 * time.Millisecond)
		fastServer := deleyedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL
		actual, err := Racer(slowURL, fastURL)
		errorChecker(err, t)
		if actual != want {
			t.Fatalf(`Want %v, but got %v`, want, actual)
		}
	})
	t.Run("should return faster", func(t *testing.T) {
		slowServer := deleyedServer(20 * time.Millisecond)
		fastServer := deleyedServer(10 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL
		actual, err := Racer(fastURL, slowURL)
		errorChecker(err, t)
		if actual != want {
			t.Fatalf(`Want %v, but got %v`, want, actual)
		}
	})

	t.Run(`should timeout`, func(t *testing.T) {
		slowServer := deleyedServer(20 * time.Second)
		fastServer := deleyedServer(15 * time.Second)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		_, err := Racer(fastURL, slowURL)
		if err == nil {
			t.Fatalf(" time out")
		}
	})

}

func deleyedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func errorChecker(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Fatal("timeout")
	}
}
