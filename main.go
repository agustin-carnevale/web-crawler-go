package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

const maxConcurrency = 15

func main() {
	// args without the program name
	args := os.Args[1:]
	if len(args) < 1 {
		// Println + exit code 1
		log.Fatalln("no website provided")
	}
	if len(args) > 1 {
		// Println + exit code 1
		log.Fatalln("too many arguments provided")
	}

	baseUrlStr := args[0]
	baseURL, err := url.Parse(baseUrlStr)
	if err != nil {
		log.Fatalln("couldn't parse url:", baseUrlStr)
	}

	fmt.Println("starting crawl of:", baseUrlStr)

	config := Config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	config.wg.Add(1)
	config.crawlPage(baseUrlStr)
	config.wg.Wait() // Ensure all goroutines complete before exiting

	for key, value := range config.pages {
		fmt.Printf("%s: %d\n", key, value)
	}
}
