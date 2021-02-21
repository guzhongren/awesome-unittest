package iteration

import "testing"

func TestRepeat(t *testing.T) {
	actual := Repeat(`a`, 6)
	expect := `aaaaaa`
	if actual != expect {
		t.Errorf(`expect %s, but got %s`, expect, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(`b`, 5)
	}
}
