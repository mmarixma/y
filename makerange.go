package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	list := make([]int, max-min)
	for i := min; i < max; i++ {
		list[i-min] = i
	}
	return list
}
