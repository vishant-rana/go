package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var str string
	var arguments = os.Args
	if len(arguments) == 1 {
		fmt.Print("put at least one argument")
	} else {
		str = arguments[1]
	}
	io.WriteString(os.Stdout, "this is standard output")
	io.WriteString(os.Stderr, str)
	io.WriteString(os.Stderr, "\n")

}
