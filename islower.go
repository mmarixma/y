package piscine

func IsLower(s string) bool {
	for _, r := range s {
		lower := r >= 'a' && r <= 'z'
		if !lower {
			return false
		}
	}
	return true
}
