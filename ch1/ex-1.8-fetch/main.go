// fetch prints the content found at a URL (like curl)
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]

	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	fmt.Printf("\n\nMaking GET request to %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("HTTP status: %d\n\n", resp.StatusCode)

	bytesWritten, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Printf("Wrote %d bytes\n", bytesWritten)
}
