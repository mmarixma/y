package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	list := make([]int, max-min)
	for i := 0; i < len(list); i++ {
		list[i] = min + i
	}
	return list
}
