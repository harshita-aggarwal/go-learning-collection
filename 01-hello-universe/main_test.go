package main

import "testing"

func TestGreet (t *testing.T) {

	got := Greet()
	want := "Hello, Universe!"

	if got != want {
		t.Errorf("Greet() printed %q; want %q", got, want)
	}
}