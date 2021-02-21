package main

import "strconv"

// FizzBuzz 类
type FizzBuzz struct{}

// Get 根据输入值获取转换后的值
func (fb FizzBuzz) Get(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "fizzbuzz"
	}
	if number%3 == 0 {
		return "fizz"
	}
	if number%5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(number)
}
