package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb <= 1 {
		return 1
	}
	prev := RecursiveFactorial(nb - 1)
	result := nb * prev
	if result/nb != prev {
		return 0
	}
	return result
}
