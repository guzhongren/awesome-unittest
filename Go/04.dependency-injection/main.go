package main

import (
	"fmt"
	"io"
	"os"
)

// Greet 打印问候
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, "+name)
}

func main() {
	Greet(os.Stdout, "test")
}
