package main

import (
	"os"

	"github.com/01-edu/z01"
)

func convertToInt(s string) int {
	digit := 0
	for _, char := range s {
		digit = (digit * 10) + int(char-'0')
	}
	return digit
}

func main() {
	userArguments := os.Args[1:]
	if len(userArguments) < 1 {
		return
	}
	baseStartRuneSmall := 'a'
	baseStartRuneCapital := 'A'

	for _, val := range userArguments {
		num := convertToInt(val)
		if val == "--upper" {
			continue
		}
		if userArguments[0] == "--upper" {
			if num >= 1 && num <= 26 {
				z01.PrintRune(baseStartRuneCapital + rune(num-1))
			} else {
				z01.PrintRune(' ')
			}
		} else {
			if num >= 1 && num <= 26 {
				z01.PrintRune(baseStartRuneSmall + rune(num-1))
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
