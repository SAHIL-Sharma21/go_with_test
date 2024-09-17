package main

import "fmt"

func main() {
	fmt.Println("Gol lang with testing")
	fmt.Println("Hello, World!")
	fmt.Println(HelloUser("Manish"))
}

const (
	englishHelloPrefix = "Hello, "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
)

func Hello() string {
	return "Hello Golang, with Testing"
}

func HelloUser(username string) string {
	if username == "" {
		username = "World!"
	}
	return englishHelloPrefix + username
}

func GreetUser(username, language string) string {
	if username == "" {
		username = "World!"
	}

	return greetingPrefix(language) + username
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
