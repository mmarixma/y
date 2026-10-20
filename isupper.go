package piscine

func IsUpper(s string) bool {
	for _, r := range s {
		upper := r >= 'A' && r <= 'Z'
		if !upper {
			return false
		}
	}
	return true
}
