package main

import "fmt"

var helloPrefix = "Hello, "

func getHelloPrefix(language string) (prefix string) {
	switch language {
	case `chinese`:
		prefix = "你好, "
	case `spanish`:
		prefix = `Hola, `
	default:
		prefix = `Hello, `
	}
	return
}

// Hello output Hello to someone
func Hello(language, name string) string {
	if name == "" {
		name = "world"
	}
	return getHelloPrefix(language) + name
}
func main() {
	fmt.Println(Hello(``, `li`))
}
