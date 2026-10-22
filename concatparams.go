package piscine

func ConcatParams(args []string) string {
	var result string
	for i, r := range args {
		result += r
		if i != len(args)-1 {
			result += "\n"
		}
	}
	return result
}
