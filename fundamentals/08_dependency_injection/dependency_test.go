package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sahil")

	got := buffer.String()
	want := "Hello, Sahil"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
