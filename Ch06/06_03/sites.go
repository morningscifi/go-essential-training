// Get content type of sites (using channels)
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s -> error: %s\n", url, err)
		return
	}
	// Should quiz myself on the difference between:
	// Sprintf, Printf, Fscanf, etc.
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s", url, ctype)
	ch <- ctype
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// Create response channel
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}

	// TODO: Wait using channel
	for i := range urls {
		output := <-ch
		fmt.Printf("\n%d Output: %s\n", i, output)
	}
}
