package piscine

func ReverseMenuIndex(menu []string) []string {
	answer := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		answer[i] = menu[len(menu)-1-i]
	}
	return answer
}
