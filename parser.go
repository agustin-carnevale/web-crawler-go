package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error getting html from %s, status code: %d", rawURL, resp.StatusCode)
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("error getting html from %s, invalid content-type: %s", rawURL, resp.Header.Get("Content-Type"))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

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
