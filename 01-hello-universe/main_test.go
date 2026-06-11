package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		lang string
		want string
	}{
		{"en", "Hello, Universe!"},
		{"fr", "Bonjour, Univers!"},
		{"de", "Hallo, Universum!"},
		{"id", "Halo, Semesta!"},
		{"it", "Ciao, Universo!"},
		{"sv", "Hej, Universum!"},
	}

	for _, test := range tests {
		got := Greet(test.lang)

		if got != test.want {
			t.Errorf("Greet(%q) = %q; want %q",
				test.lang, got, test.want)
		}
	}
}
