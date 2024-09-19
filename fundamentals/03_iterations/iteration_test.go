package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	assertCorrectMessage(t, repeated, expected)
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("Expected %q and got %q\n", expected, repeated)
	}
}

func assertCorrectNumber(t testing.TB, repeated, expected int) {
	t.Helper()

	if repeated != expected {
		t.Errorf("Expected %d and got %d\n", expected, repeated)
	}
}

// benchmark test
// Writing benchmarks in Go is another first-class feature of the language and it is very similar to writing tests.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestRepeatWithCount(t *testing.T) {
	repeated := RepeatWithCount("a", 2)
	expected := "aa"
	assertCorrectMessage(t, repeated, expected)
}

// Example fun call
func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	//output: aaaaa
}

func TestSumNumber(t *testing.T) {
	t.Run("sum of number which passed", func(t *testing.T) {
		repeated := SumNumber(5)
		expected := 10
		assertCorrectNumber(t, repeated, expected)
	})
	t.Run("test when no parameter is passed", func(t *testing.T) {
		repeated := SumNumber(0)
		expected := 0
		assertCorrectNumber(t, repeated, expected)
	})
}

func TestIsContains(t *testing.T) {
	t.Run("is given username contains", func(t *testing.T) {
		got := IsContains("Sahil", "i")
		expected := true
		if got != expected {
			t.Errorf("Expected %t and got %t\n", expected, got)
		}
	})
	t.Run("given username is not pressent", func(t *testing.T) {
		got := IsContains("", "i")
		expected := false
		if got != expected {
			t.Errorf("Expected %t and got %t\n", expected, got)
		}
	})
}
