package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run(`test output in sub test`, func(t *testing.T) {
		buffer := bytes.Buffer{}
		Countdown(&buffer)

		got := buffer.String()
		want := `3
2
1
go`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
