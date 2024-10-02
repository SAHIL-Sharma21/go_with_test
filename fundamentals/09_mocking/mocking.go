package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("mocking in go with test")
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const countDownStart = 3

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

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
