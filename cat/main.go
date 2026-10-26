package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	failed := false
	arrgs := os.Args[1:]

	if len(arrgs) == 0 {
		printContent(os.Stdin)
		return
	}

	for _, name := range arrgs {
		file, err := os.Open(name)
		if err != nil {
			printStr("ERROR: " + err.Error() + "\n")
			failed = true
			continue
		}
		printContent(file)
		file.Close()
	}
	if failed {
		os.Exit(1)
	}
}

func printContent(file *os.File) {
	arr := make([]byte, 100)

	for {
		n, err := file.Read(arr)
		if n > 0 {
			for _, b := range arr[:n] {
				z01.PrintRune(rune(b))
			}
		}
		if err != nil {
			break
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
