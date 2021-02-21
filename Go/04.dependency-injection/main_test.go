package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// 申明 buffer，准备接受数据， 因为bytes.Buffer， 重点：bytes.Buffer实现了 io.Writer
	buffer := bytes.Buffer{}
	// 将buffer 传入，此时就是依赖注入的入口，
	Greet(&buffer, "chris")
	// 获取程序运行的结果
	got := buffer.String()
	// 期望值
	want := "Hello, chris"
	// 测试判断
	if got != want {
		t.Errorf(`got %s, want %s`, got, want)
	}
}
