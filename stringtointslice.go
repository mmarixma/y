package piscine

func StringToIntSlice(str string) []int {
	var answer []int
	for _, r := range str {
		answer = append(answer, int(r))
	}
	return answer
}
