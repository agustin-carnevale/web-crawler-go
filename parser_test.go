package main

import (
	"reflect"
	"testing"
)

func TestGetUrlsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
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
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs and some nesting",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<div>
				<p>
					This is another link 
					<a href="https://boot.dev">
						<span>Boot.dev</span>
					</a>
				</p>
			</div>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://boot.dev", "https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no a tags",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<div>
						<p>
							This is another link 
								<span>Boot.dev</span>
						</p>
					</div>
						<span>Hey there!</span>
					<p>
						<span>Boot.dev</span>
					</p>
				</body>
			</html>
			`,
			expected: []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: \nexpected: %v \nactual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
