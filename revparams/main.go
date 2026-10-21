package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		name := os.Args[i]
		for _, r := range name {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
