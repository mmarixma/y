package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {
	userArguments := os.Args[1:]
	hasInsert := false
	hasOrder := false
	insertVal := ""
	targetVal := ""

	if len(userArguments) == 0 {
		printHelp()
		return
	}

	for i := 0; i < len(userArguments); i++ {
		if userArguments[i] == "--help" || userArguments[i] == "-h" {
			printHelp()
			return
		}
		if userArguments[i] == "--order" || userArguments[i] == "-o" {
			hasOrder = true
		}
		if userArguments[i] == "--insert" || userArguments[i] == "-i" {
			hasInsert = true
			insertVal = userArguments[i+1]
			i++
		}
		if len(userArguments[i]) >= 9 && userArguments[i][0:9] == "--insert=" {
			hasInsert = true
			insertVal = userArguments[i][9:]
		}
		if len(userArguments[i]) >= 3 && userArguments[i][0:3] == "-i=" {
			hasInsert = true
			insertVal = userArguments[i][3:]
		}
		if len(userArguments[i]) == 0 || userArguments[i][0] != '-' {
			targetVal = userArguments[i]
		}
	}

	if hasInsert {
		targetVal += insertVal
	}

	targetValArr := []rune(targetVal)

	if hasOrder {
		for i := 0; i < len(targetValArr); i++ {
			for j := 0; j < len(targetValArr)-i-1; j++ {
				if targetValArr[j] > targetValArr[j+1] {
					targetValArr[j], targetValArr[j+1] = targetValArr[j+1], targetValArr[j]
				}
			}
		}
	}
	fmt.Println(string(targetValArr))
}
