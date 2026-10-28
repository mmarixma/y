package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := true
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				for l := '9'; l >= '0'; l-- {
					if i < k || (i == k && j <= l) {
						continue
					}
					if !first {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					first = false
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)
				}
			}
		}
	}
}
