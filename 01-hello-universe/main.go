package main

import (
	"fmt"
)

func Greet() string {
	return "Hello, Universe!"
}

func main() {
	s := Greet()
	fmt.Println(s)
}