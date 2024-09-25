package main

import (
	"reflect"
	"testing"
)

// func TestSum(t *testing.T) {
// 	numbers := [5]int{2, 4, 6, 8, 10}

// 	got := Sum(numbers)
// 	expected := 0

// 	assertMessage(t, got, expected, numbers[:])
// }

func TestSumNumber(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := SumNumbers(numbers[:])
	expected := 15

	assertMessage(t, got, expected, numbers[:])
}

func assertMessage(t testing.TB, got int, expected int, given []int) {
	t.Helper()

	if got != expected {
		t.Errorf("got %d want %d given %v", got, expected, given)
	}

}

// writing test for array and slice
func TestSum(t *testing.T) {

	t.Run("test with limited nummbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers[:])
		expected := 15

		assertMessage(t, got, expected, numbers[:])
	})

	t.Run("test with any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		assertMessage(t, got, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if reflect.DeepEqual(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}

// Sum all tails
func TestSumAllTails(t *testing.T) {

	//refactoring of the code
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	}

	t.Run("slices with not empty input", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, got, expected)
	})

	t.Run("slices with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3})
		expected := []int{0, 3}

		checkSums(t, got, expected)
	})
}
