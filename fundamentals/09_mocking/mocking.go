package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("mocking in go with test")
	// sleeper := &DefaultSleeper{}
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const countDownStart = 3

/*
func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
*/
// real world mocking
type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// func CountDown(out io.Writer, sleeper Sleeper) {
// 	for i := countDownStart; i > 0; i-- {
// 		sleeper.Sleep()
// 	}

// 	for i := countDownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 	}

// 	sleeper.Sleep()
// 	fmt.Fprint(out, finalWord)
// }

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// Extending Sleeper to be configurable
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
