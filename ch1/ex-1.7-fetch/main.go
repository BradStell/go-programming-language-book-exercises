// fetch prints the content found at a URL (like curl)
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	bytesWritten, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Printf("Wrote %d bytes\n", bytesWritten)
}