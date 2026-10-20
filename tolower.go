package piscine

func ToLower(s string) string {
	lower := []rune(s)
	for i, r := range lower {
		if r >= 'A' && r <= 'Z' {
			lower[i] = r + 32
		}
	}
	return string(lower)
}
