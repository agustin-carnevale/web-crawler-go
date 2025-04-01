package main

import (
	"fmt"
	"sort"
	"strings"
)

type reportPageItem struct {
	numOfLinks int
	page       string
}

func printReport(pages map[string]int, baseURL string) {

	fmt.Println("=============================")
	fmt.Println(" REPORT for", baseURL)
	fmt.Println("=============================")
	fmt.Println("")

	pagesList := sortPages(pages)

	for _, p := range pagesList {
		fmt.Printf("Found %d internal links to %s\n", p.numOfLinks, p.page)
	}
	fmt.Println("")
}

func sortPages(pages map[string]int) []reportPageItem {
	results := []reportPageItem{}

	for page, n := range pages {
		results = append(results, reportPageItem{
			page:       page,
			numOfLinks: n,
		})
	}

	// Sort by numOfLinks in desc order
	sort.Slice(results, func(i, j int) bool {
		if results[i].numOfLinks != results[j].numOfLinks {
			return results[i].numOfLinks > results[j].numOfLinks
		}

		return strings.Compare(results[i].page, results[j].page) < 0
	})

	return results
}
