package piscine

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	result := []rune(s)
	startOfWord := true
	for i, c := range result {
		if isAlphanumeric(c) {
			if startOfWord {
				result[i] = toUpper(c)
				startOfWord = false
			} else {
				result[i] = toLower(c)
			}
		} else {
			startOfWord = true
		}
	}
	return string(result)
}

func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toUpper(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

func toLower(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}