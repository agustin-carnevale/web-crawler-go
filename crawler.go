package main

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *Config) crawlPage(rawCurrentURL string) {
	fmt.Println("")
	fmt.Println("Current Page:", rawCurrentURL)
	fmt.Println("")

	// take 1 space in the buffer channel
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		// free 1 space in the buffer channel
		<-cfg.concurrencyControl
		// Decrement wg
		cfg.wg.Done()
	}()

	currentUrl, _ := url.Parse(rawCurrentURL)

	if normalizeHost(cfg.baseURL.Host) != normalizeHost(currentUrl.Host) {
		// we only care to parse pages within the same domain
		return
	}

	normalizedCurrentUrl, _ := normalizeURL(rawCurrentURL)

	// Thread safe access to pages map
	cfg.mu.Lock()
	if _, exists := cfg.pages[normalizedCurrentUrl]; exists {
		cfg.pages[normalizedCurrentUrl] += 1
		cfg.mu.Unlock()
		return
	}
	cfg.pages[normalizedCurrentUrl] = 1
	cfg.mu.Unlock()
	//*********

	currentHtml, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("*** CURRENT PAGE HTML ***")
	// fmt.Println(currentHtml)

	links, err := getURLsFromHTML(currentHtml, cfg.baseURL.String())
	if err != nil {
		return
	}
	// fmt.Println("*** CURRENT PAGE LINKS ***")
	// fmt.Printf("%v\n", links)

	for _, href := range links {
		// Increment wg before launching a goroutine
		cfg.wg.Add(1)
		go cfg.crawlPage(href)
	}
}
