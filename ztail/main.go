package main

import (
	"fmt"
	"os"
)

func stringToInt(str string) int {
	digit := 0

	for _, char := range str {
		digit = (digit * 10) + int(char-'0')
	}
	return digit
}

func main() {
	userArguments := os.Args[1:]

	tail := 0
	exitCode := 0
	start := 0

	for i := 0; i < len(userArguments); i++ {
		if userArguments[i] == "-c" {
			tail = stringToInt(userArguments[i+1])
			i++
			continue
		}
		file, err := os.Open(userArguments[i])
		if err != nil {
			exitCode = 1
			fmt.Println(err.Error())
			continue
		}
		arr := make([]byte, 200)
		n, _ := file.Read(arr)
		if len(userArguments) > 3 {
			if i != 2 {
				fmt.Println()
			}

			fmt.Println("==> " + userArguments[i] + " <==")
		}
		if tail < n {
			start = n - tail
		}
		fmt.Print(string(arr[start:n]))
		file.Close()
	}
	if exitCode != 0 {
		os.Exit(1)
	}
}
