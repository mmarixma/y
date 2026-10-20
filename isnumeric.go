package piscine

func IsNumeric(s string) bool {
	for _, r := range s {
		isDigit := r >= '0' && r <= '9'
		if !isDigit {
			return false
		}
	}
	return true
}
