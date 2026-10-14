package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	loop(n)
}

func loop(n int) {
	if n >= 10 || n <= -10 {
		loop(n / 10)
	}
	d := n % 10
	if d < 0 {
		d = -d
	}
	z01.PrintRune(rune(d + '0'))
}
