package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// fmt.Println(normalizeURL("https://blog.boot.dev/path"))

	inputURL := "https://blog.boot.dev"
	inputBody := `
<html>
<body>
	<a href="/path/one">
		<span>Boot.dev</span>
	</a>
	<a href="https://other.com/path/one">
		<span>Boot.dev</span>
	</a>
</body>
</html>
`

	getURLsFromHTML(inputBody, inputURL)
}
