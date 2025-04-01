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

	// 	// fmt.Println(normalizeURL("https://blog.boot.dev/path"))

	// 	inputURL := "https://blog.boot.dev"
	// 	inputBody := `
	// <html>
	// <body>
	// 	<a href="/path/one">
	// 		<span>Boot.dev</span>
	// 	</a>
	// 	<a href="https://other.com/path/one">
	// 		<span>Boot.dev</span>
	// 	</a>
	// </body>
	// </html>
	// `

	// getURLsFromHTML(inputBody, inputURL)
}
