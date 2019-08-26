package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Add all arguments to an array of filenames. If no arguments are given, add "-" (for stdin).
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"-"}
	}

	// Read each file line by line and print the contents.
	for _, fileName := range args {
		func() {
			var file *os.File
			var fileErr error
			if fileName == "-" {
				file = os.Stdin
			} else {
				file, fileErr = os.Open(fileName)
				if fileErr != nil {
					log.Fatal(fileErr)
				}
				defer func() {
					if err := file.Close(); err != nil {
						log.Fatal(err)
					}
				}()
			}

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			if scanErr := scanner.Err(); scanErr != nil {
				if _, err := fmt.Fprintf(os.Stderr, "reading file %s: %s", fileName, scanErr); err != nil {
					// If we cannot write to stderr, something is very wrong.
					panic(err)
				}
			}

		}()
	}
}
