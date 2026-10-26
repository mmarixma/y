package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		input := make([]byte, 200)
		for {
			n, err := os.Stdin.Read(input)
			printError(string(input[:n]))
			if err == io.EOF {
				break
			}
		}
	}
	for _, file := range args {
		file, err := os.Open(file)
		if err != nil {
			printError("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		arr := make([]byte, 100)
		n, err := file.Read(arr)
		printError(string(arr[:n]))
		file.Close()
	}
}

func printError(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
