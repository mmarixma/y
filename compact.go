package piscine

func Compact(ptr *[]string) int {
	write := 0
	for _, r := range *ptr {
		if r != "" {
			(*ptr)[write] = r
			write++
		}
	}
	*ptr = (*ptr)[:write]
	return write
}
