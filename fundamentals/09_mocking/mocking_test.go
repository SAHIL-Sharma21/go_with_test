package main

import (
	"bytes"
	"testing"
)

// func TestCountDown(t *testing.T) {
// 	buffer := &bytes.Buffer{}

// 	CountDown(buffer)

// 	got := buffer.String()
// 	want := `3
// 2
// 1
// Go!`

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// real mocking
func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleep := &SpySleeper{}

	CountDown(buffer, spySleep)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleep.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleep.Calls)
	}
}
