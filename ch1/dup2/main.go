package main

import (
	"bufio"
	"fmt"
	"os"
)

type result struct {
	hits  int
	files map[string]bool
}

func main() {
	counts := make(map[string]*result)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.hits > 1 {
			fileString := ""
			for key := range n.files {
				fileString += key + " "
			}
			fmt.Printf("%d\t%s\t%s\n", n.hits, line, fileString)
		}
	}
}

func countLines(f *os.File, counts map[string]*result) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = &result{0, make(map[string]bool)}
		}
		counts[input.Text()].hits++
		counts[input.Text()].files[f.Name()] = true
	}
}
