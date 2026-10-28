package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := map[string]int{}
	if str == "" {
		list[str] += 1
		return list
	}
	newString := ""
	for idx, char := range str {
		if char == ' ' {
			list[newString] += 1
			newString = ""
			continue
		}

		if idx == len(str)-1 {
			newString += string(char)
			list[newString] += 1
			continue
		}
		newString += string(char)
	}
	if str[len(str)-1] == ' ' {
		list[""] += 1
	}
	return list
}
