package main

import (
	"fmt"
	"strings"
)

/*-----------------*/
/*Greet RETURNS GREETINGS IN THE LANGUAGE CODE PROVIDED AS INPUT*/
/*-----------------*/
func Greet(l string) string {

	l = strings.ToLower(l)

	langs := map[string]string{
		"en": "Hello, Universe!",
		"id": "Halo, Semesta!",
		"fr": "Bonjour, Univers!",
		"de": "Hallo, Universum!",
		"it": "Ciao, Universo!",
		"sv": "Hej, Universum!",
	}

	greeting, exists := langs[l]
	if exists {
		return greeting
	} else {
		return fmt.Sprintf("I don't know the language code: %s", l)
	}
}

func main() {

	var l string

	fmt.Print("Enter language code: ")
	fmt.Scanln(&l)

	fmt.Println(Greet(l))
}
