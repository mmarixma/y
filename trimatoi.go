package piscine

func TrimAtoi(s string) int {
	if s == "" {
		return 0
	}
	i := 0
	neg := false
	found := false
	n := 0

	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			found = true
			n = n*10 + int(s[i]-'0')
		} else if s[i] == '-' && !found {
			neg = true
		}
		i++
	}

	if neg {
		n = -n
	}
	return n
}
