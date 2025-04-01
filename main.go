package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

const defaultMaxConcurrency = 15
const defaultMaxPages = 100

func main() {
	// usage: ./crawler URL [maxConcurrency] [maxPages]

	// args without the program name
	args := os.Args[1:]
	if len(args) < 1 {
		// Println + exit code 1
		log.Fatalln("no website provided")
	}
	if len(args) > 3 {
		// Println + exit code 1
		log.Fatalln("too many arguments provided, usage: ./crawler URL [maxConcurrency] [maxPages]")
	}

	baseUrlStr := args[0]
	baseURL, err := url.Parse(baseUrlStr)
	if err != nil {
		log.Fatalln("couldn't parse url:", baseUrlStr)
	}

	maxConcurrency := defaultMaxConcurrency
	maxPages := defaultMaxPages

	if len(args) > 1 {
		maxConcurrencyArg, err := strconv.Atoi(args[1])
		if err == nil {
			maxConcurrency = maxConcurrencyArg
		}
	}
	if len(args) > 2 {
		maxPagesArg, err := strconv.Atoi(args[2])
		if err == nil {
			maxPages = maxPagesArg
		}

	}

	fmt.Println("starting crawl of:", baseUrlStr)

	config := Config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	config.wg.Add(1)
	config.crawlPage(baseUrlStr)
	config.wg.Wait() // Ensure all goroutines complete before exiting

	for key, value := range config.pages {
		fmt.Printf("%s: %d\n", key, value)
	}
}
