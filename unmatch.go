package piscine

func Unmatch(a []int) int {
	count := 0
	for _, val1 := range a {
		for _, val2 := range a {
			if val1 == val2 {
				count++
			}
		}
		if count%2 != 0 {
			return val1
		}
		count = 0
	}
	return -1
}
