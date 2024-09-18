package main

import "fmt"

func main() {
	fmt.Println("Golang with testing inintegers")
	ans := Subtract(5, 2)
	fmt.Println(ans)
}

func Adder(num1, num2 int) int {
	return num1 + num2
}

func Subtract(num1, num2 int) int {
	if num2 > num1 {
		return num2 - num1
	}
	return num1 - num2
}
