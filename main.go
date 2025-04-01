package main

import (
	"fmt"
	"log"
	"os"
)

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

	baseUrl := args[0]

	fmt.Println("starting crawl of:", baseUrl)

	// rawHtml, _ := getHTML(baseUrl)
	// fmt.Println(rawHtml)

	pages := make(map[string]int)
	crawlPage(baseUrl, baseUrl, pages)

	for key, value := range pages {
		fmt.Printf("%s: %d\n", key, value)
	}
}
