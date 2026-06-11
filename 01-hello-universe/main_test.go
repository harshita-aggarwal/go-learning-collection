package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		lang string
		want string
	}{
		{"English", "en", "Hello, Universe!"},
		{"French", "fr", "Bonjour, Univers!"},
		{"German", "de", "Hallo, Universum!"},
		{"Indonesian", "id", "Halo, Semesta!"},
		{"Italian", "it", "Ciao, Universo!"},
		{"Swedish", "sv", "Hej, Universum!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := Greet(test.lang)

			if got != test.want {
				t.Errorf("Greet(%q) = %q; want %q",
					test.lang, got, test.want)
			}
		})
	}
}
