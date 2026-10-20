package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		isChar := r >= ' ' && r <= '~'
		if !isChar {
			return false
		}
	}
	return true
}
