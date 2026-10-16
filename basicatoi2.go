package piscine

func BasicAtoi2(s string) int {
	for _, r := range s {
		up := r >= 'A' && r <= 'Z'
		low := r >= 'a' && r <= 'z'
		if up || low {
			return 0
		}
		if r == ' ' {
			return 0
		}
	}
	n := 0
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
