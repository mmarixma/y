package piscine

func JumpOver(str string) string {
	if str == "" || len(str) < 3 {
		return "\n"
	}
	newString := ""
	for i := 2; i < len(str); i = i + 3 {
		newString += string(str[i])
	}
	newString += "\n"
	return newString
}
