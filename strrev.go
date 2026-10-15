package piscine

func StrRev(s string) string {
	// results := ""
	// arr := []rune(s)
	// for i := len(arr) - 1; i >= 0; i-- {
	// 	results += string(arr[i])
	// }
	// return results

	result := ""
	for _, r := range s {
		result = string(r) + result
	}
	return result
}
