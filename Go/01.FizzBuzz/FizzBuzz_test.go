package main

import "testing"

func TestFizzBussInstance(test *testing.T) {
	fizzBuzz := FizzBuzz{}
	_, ok := interface{}(fizzBuzz).(FizzBuzz)
	if !ok {
		test.Errorf("FizzBuzzTest.go instance: expect %v, but got %v", FizzBuzz{}, fizzBuzz)
	}
}
func TestGetFizzBuzz(t *testing.T) {
	fizzBuzz := FizzBuzz{}
	tests := []struct {
		input  int
		output string
	}{
		{1, "1"},
		{3, "fizz"},
		{5, "buzz"},
		{15, "fizzbuzz"},
		{100, "buzz"},
	}
	for _, tt := range tests {
		actual := fizzBuzz.Get(tt.input)
		if actual != tt.output {
			t.Errorf("FizzBuzz.Get(%d), expect %s, but got %s", tt.input, tt.output, actual)
		}
	}
}
