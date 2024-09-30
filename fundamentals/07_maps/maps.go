package main

import "fmt"

func main() {

	fmt.Println("maps with test")
}

// first serach function
func Search(dict map[string]string, word string) string {
	return dict[word]
}
