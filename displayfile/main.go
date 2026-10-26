package main

import (
	"fmt"
	"os"
)

func main() {
	arrgs := os.Args[1:]

	if len(arrgs) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(arrgs) != 1 {
		fmt.Println("Too many arguments")
		return
	}

	file, _ := os.Open(arrgs[0])

	arr := make([]byte, 100)

	n, _ := file.Read(arr)

	fmt.Print(string(arr[:n]))
	file.Close()
}
