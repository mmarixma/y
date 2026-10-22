package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var result string
	for i, r := range a {
		result += r
		if i != len(a)-1 {
			result += "\n"
		}
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
