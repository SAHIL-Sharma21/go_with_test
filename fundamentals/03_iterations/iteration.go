package main

import (
	"fmt"
	"strings"
)

const (
	repeatedCount = 5
)

func main() {
	fmt.Println("Itertaion in go lang with test")

	// ans := Repeat("a")
	// fmt.Println("ans:", ans)

	// sum := SumNumber(3)
	// fmt.Println(sum)

	isTrue := IsContains("sahil", "h")
	fmt.Println("ispresent:", isTrue)
}

// func Repeat(character string) string {
// 	return ""
// }

func Repeat(character string) string {
	out := ""
	for i := 0; i < repeatedCount; i++ {
		out += character
	}
	return out
}

// excersise
func RepeatWithCount(character string, count int) string {
	out := ""
	for i := 0; i < count; i++ {
		out += character
	}
	return out
}

func SumNumber(num int) int {
	sum := 0

	for i := 0; i < num; i++ {
		sum += i
	}
	return sum
}

// with strings
// function to check if the substr is present or not
func IsContains(username, substr string) bool {
	return strings.Contains(username, substr)
}
