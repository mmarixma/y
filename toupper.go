package piscine

func ToUpper(s string) string {
	cap := []rune(s)
	for i, r := range cap {
		low := r >= 'a' && r <= 'z'
		if low {
			cap[i] = r - 32
		}
	}
	return string(cap)
}
