package main

import "testing"

// card reader that takes a []string and fires off 
// a goroutine to copy it to a string channel
func dataReader(data []string) chan string {
	cr := make(chan string)

	go func() {
		for _, card := range data {
			cr <- card
		}
		close(cr)
	}()

	return cr
}

// Just a simple way to take everything out of a 
// string channel used for benchmarking
func consumer(feed chan string) {
	for {
		_, ok := <-feed
		if !ok {
			return
		}
	}
}

