// a programm to search files for stuff
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: content_search <query>")
		return
	}

	query := strings.Join(os.Args[1:], " ")
	query_colored := "\033[31;1m" + query + "\033[0m" // the same but red

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
			return err
		}

		line := 0
		scanner := bufio.NewScanner(f)
		first := true

		for scanner.Scan() {
			if strings.Contains(scanner.Text(), query) {
				if first {
					fmt.Printf("%s\n", path)
					first = false
				}
				fmt.Printf(":%d %s\n", line, strings.Replace(scanner.Text(), query, query_colored, -1))
			}
			line++
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
	}
}
