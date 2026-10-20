package piscine

func IsAlpha(s string) bool {
	for _, r := range s {
		upper := r >= 'A' && r <= 'Z'
		lower := r >= 'a' && r <= 'z'
		isDigit := r >= '0' && r <= '9'
		if !upper && !lower && !isDigit {
			return false
		}
	}
	return true
}
