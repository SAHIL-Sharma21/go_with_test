package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
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

/*
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
*/

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		CountDown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

// configurable test
func TestConfigurableSleeper(t *testing.T) {
	sleeptime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleeptime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleeptime {
		t.Errorf("should have slept for %v but slept for %v", sleeptime, spyTime.durationSlept)
	}
}
