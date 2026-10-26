package main

import (
	"fmt"
	"os"
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
			fmt.Println("ERROR:", err)
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
			fmt.Print(string(arr[:n]))
		}
		if err != nil {
			break
		}
	}
}
