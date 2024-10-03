package main

import (
	"fmt"
)

func main() {
	fmt.Println("Concurrency with test")
}

// website checker funtion which will check the status of the listy of URLS
type WebsiteChecker func(string) bool

// using channels
type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		//go routine
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
