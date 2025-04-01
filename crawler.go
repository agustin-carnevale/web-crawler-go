package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// fmt.Println("")
	// fmt.Println("Current Page:", rawCurrentURL)
	// fmt.Println("")

	baseUrl, _ := url.Parse(rawBaseURL)
	currentUrl, _ := url.Parse(rawCurrentURL)

	if normalizeHost(baseUrl.Host) != normalizeHost(currentUrl.Host) {
		// we only care to parse pages within the same domain
		return
	}

	normalizedCurrentUrl, _ := normalizeURL(rawCurrentURL)
	if _, exists := pages[normalizedCurrentUrl]; exists {
		pages[normalizedCurrentUrl] += 1
		return
	}

	pages[normalizedCurrentUrl] = 1

	currentHtml, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("*** CURRENT PAGE HTML ***")
	// fmt.Println(currentHtml)

	links, err := getURLsFromHTML(currentHtml, rawBaseURL)
	if err != nil {
		return
	}
	// fmt.Println("*** CURRENT PAGE LINKS ***")
	// fmt.Printf("%v\n", links)

	for _, href := range links {
		crawlPage(rawBaseURL, href, pages)
	}

}
