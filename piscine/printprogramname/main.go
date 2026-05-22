package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programPath := os.Args[0]

	shortName := ""
	for i := len(programPath) - 1; i >= 0; i-- {
		if programPath[i] == '/' {
			shortName = programPath[i+1:]
			break
		}
	}

	if shortName == "" {
		shortName = programPath
	}

	for _, char := range shortName {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
