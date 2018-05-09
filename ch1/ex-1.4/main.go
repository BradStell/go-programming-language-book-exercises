// dub2 prints the count and text of lines that appear more than once
// in the input. It reads stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	
}

func countLines(f *os.File, counts map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	printedFileName := false
	for line, n := range counts {
		if n > 1 {
			if !printedFileName {
				fmt.Printf("File Name: %s\n", fileName)
				printedFileName = true
			}

			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
