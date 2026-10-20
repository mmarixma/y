package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	num := []int{}
	for n > 0 {
		num = append(num, n%10)
		n /= 10
	}
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-1-i; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	for _, r := range num {
		z01.PrintRune(rune(r + '0'))
	}
}
