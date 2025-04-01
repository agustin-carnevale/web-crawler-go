package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

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
					link := attr.Val
					u, _ := url.Parse(attr.Val)
					if u.Host == "" {
						link = rawBaseURL + attr.Val
					}
					linksList = append(linksList, link)
					break
				}
			}
		}
	}

	return linksList, nil
}
