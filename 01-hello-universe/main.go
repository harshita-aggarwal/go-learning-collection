package main

import (
	"fmt"
	"strings"
)

func Greet(l string) string {
	
	l = strings.ToLower(l)

	switch l {
	case "en":
		return "Hello, Universe!"
	case "id":
		return "Halo, Semesta!"
	case "fr":
		return "Bonjour, Univers!"
	case "de":
		return "Hallo, Universum!"
	default:
		return fmt.Sprintf("I don't know the language code: %s", l)
	}
}

func main() {

	var l string

	fmt.Print("Enter language code: ")
	fmt.Scanln(&l)

	fmt.Println(Greet(l))
}