package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adder(4, 2)
	expected := 6

	assertCorrectMessage(t, sum, expected)
}

func assertCorrectMessage(t testing.TB, sum, expected int) {
	t.Helper()

	if sum != expected {
		t.Errorf("Expected %d and got %d\n", sum, expected)
	}
}

func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	//output: 6
}

func TestSubtract(t *testing.T) {
	t.Run("num2 is greater than num1", func(t *testing.T) {
		sub := Subtract(2, 5)
		expected := 3
		assertCorrectMessage(t, sub, expected)
	})
	t.Run("num2 is less than num1", func(t *testing.T) {
		sub := Subtract(5, 2)
		expect := 3
		assertCorrectMessage(t, sub, expect)
	})
}

func ExampleSubtract() {
	sub := Subtract(5, 2)
	fmt.Println(sub)
	//output: 3
}
