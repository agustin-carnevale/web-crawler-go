package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", nil
	}
	normalizedUrl := u.Host + u.Path
	normalizedUrl = strings.TrimSuffix(normalizedUrl, "/")
	return normalizedUrl, nil
}
