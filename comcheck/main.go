package main

import (
	"fmt"
	"os"
)

func main() {
	userArguments := os.Args[1:]

	for _, val := range userArguments {
		if val == "01" || val == "galaxy" || val == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
