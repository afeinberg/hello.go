package main

import (
	"fmt"
	"github.com/afeinberg/hello.go"
)

func main() {
	h, err := sayhello.NewHello("World")
	if err != nil {
		panic(err.Error())
	}
	
	defer h.Dispose()
	s, err := h.SayHello()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(s)
}
