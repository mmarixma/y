package piscine

func Abort(a, b, c, d, e int) int {
	arry := []int{a, b, c, d, e}
	n := len(arry)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arry[j] > arry[j+1] {
				arry[j], arry[j+1] = arry[j+1], arry[j]
			}
		}
	}
	return arry[2]
}
