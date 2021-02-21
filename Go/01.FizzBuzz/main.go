package main

import "fmt"

func main() {
	fizzBuzz := FizzBuzz{}
	for i := 1; i <= 100; i++ {
		fmt.Print(fizzBuzz.Get(i))
	}
}
