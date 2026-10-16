package piscine

func BasicAtoi(s string) int {
	i := 0
	n := 0
	for ; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
