package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	for i := 0; i <= len(s)-len(toFind); i++ {
		j := 0
		for j = 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				break
			}
		}
		if j == len(toFind) {
			return i
		}
	}
	return -1
}
