package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == `waat://furhurterwe.geds` {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	var websites = []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	want := websites
	got := CheckWebsites(mockWebsiteChecker, websites)
	t.Run("length should be equal given", func(t *testing.T) {
		if len(want) != len(got) {
			t.Fatalf(`Wanted %v, got %v`, want, got)
		}
	})
	t.Run(`value should be equal`, func(t *testing.T) {
		expected := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}
		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Want %v, got %v", want, got)
		}
	})
}

func slowChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 1000)
	for index := range urls {
		urls[index] = "a url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowChecker, urls)
	}
}
