package main

import "testing"

func TestGreet_English (t *testing.T) {

	l := "en"

	got := Greet(l)
	want := "Hello, Universe!"

	if got != want {
		t.Errorf("Greet(%q) printed %q; want %q", l, got, want)
	}
}
func TestGreet_Indonesian (t *testing.T) {

	l := "id"

	got := Greet(l)
	want := "Halo, Semesta!"

	if got != want {
		t.Errorf("Greet(%q) printed %q; want %q", l, got, want)
	}
}
func TestGreet_French (t *testing.T) {

	l := "fr"

	got := Greet(l)
	want := "Bonjour, Univers!"

	if got != want {
		t.Errorf("Greet(%q) printed %q; want %q", l, got, want)
	}
}
func TestGreet_German (t *testing.T) {

	l := "de"

	got := Greet(l)
	want := "Hallo, Universum!"

	if got != want {
		t.Errorf("Greet(%q) printed %q; want %q", l, got, want)
	}
}
func TestGreet_Italian (t *testing.T) {

	l := "it"

	got := Greet(l)
	want := "Ciao, Universo!"

	if got != want {
		t.Errorf("Greet(%q) printed %q; want %q", l, got, want)
	}
}
func TestGreet_Swedish (t *testing.T) {

	l := "sv"

	got := Greet(l)
	want := "Hej, Universum!"

	if got != want {
		t.Errorf("Greet(%q) printed %q; want %q", l, got, want)
	}
}
func TestGreetUnknownLanguage(t *testing.T) {
	l := "es"
	got := Greet(l)
	want := "I don't know the language code: es"

	if got != want {
		t.Errorf("Greet(%q) = %q; want %q", l, got, want)
	}
}