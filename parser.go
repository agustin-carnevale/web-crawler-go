package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Couldn't parse baseUrl: %v: %v", rawBaseURL, err)

	}
	htmlReader := strings.NewReader(htmlBody)
	body, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	linksList := []string{}
	for node := range body.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					href, err := url.Parse(attr.Val)
					if err != nil {
						fmt.Printf("Couldn't parse href: %v: %v", attr.Val, err)
						continue
					}
					resolvedUrl := baseUrl.ResolveReference(href)
					linksList = append(linksList, resolvedUrl.String())
					break
				}
			}
		}
	}

	return linksList, nil
}
