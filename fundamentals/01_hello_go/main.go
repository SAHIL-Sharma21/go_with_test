package main

import "fmt"

func main() {
	fmt.Println("Gol lang with testing")
	fmt.Println("Hello, World!")
	fmt.Println(HelloUser("Manish"))
}

const englishHelloPrefix = "Hello, "

func Hello() string {
	return "Hello Golang, with Testing"
}

func HelloUser(username string) string {
	if username == "" {
		username = "World!"
	}
	return englishHelloPrefix + username
}
