package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}

	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	buffer := ""
	skip := false

	for i := 0; i < len(str); i++ {
		if skip {
			skip = false
			continue
		}

		if str[i] == ' ' {
			continue
		}

		buffer += string(str[i])

		if len(buffer) == 5 {
			if result != "" {
				result += " "
			}
			result += buffer
			buffer = ""
			skip = true
		}
	}

	if buffer != "" {
		if result != "" {
			result += " "
		}
		result += buffer
	}

	return result + "\n"
}
