package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}
	for _, r := range a {
		result = append(result, f(r))
	}
	return result
}
