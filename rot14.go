package piscine

func Rot14(s string) string {
	arry := []rune(s)
	for i := 0; i < len(arry); i++ {
		if arry[i] >= 'a' && arry[i] <= 'z' {
			arry[i] = arry[i] + 14
			if arry[i] > 'z' {
				arry[i] = arry[i] - 26
			}
			continue
		}
		if arry[i] >= 'A' && arry[i] <= 'Z' {
			arry[i] = arry[i] + 14
			if arry[i] > 'Z' {
				arry[i] = arry[i] - 26
			}
		}
	}
	return string(arry)
}
