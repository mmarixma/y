package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	list := []int{}
	for i := min; i < max; i++ {
		list = append(list, i)
	}
	return list
}
