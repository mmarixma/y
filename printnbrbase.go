package piscine

import "github.com/01-edu/z01"

func printNV() {
	z01.PrintRune('N')
	z01.PrintRune('V')
}

func recurseBase(nbr uint64, base string) {
	if nbr >= uint64(len(base)) {
		recurseBase(nbr/uint64(len(base)), base)
	}
	z01.PrintRune(rune(base[nbr%uint64(len(base))]))
}

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		printNV()
		return
	}
	for i := 0; i < len(base); i++ {
		for j := 0; j < len(base); j++ {
			if (i != j && base[i] == base[j]) || base[i] == '-' || base[i] == '+' {
				printNV()
				return
			}
		}
	}
	var number uint64
	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			number = 9223372036854775808
		} else {
			number = uint64(-nbr)
		}
	} else {
		number = uint64(nbr)
	}
	recurseBase(number, base)
}
