package piscine

func IsPrime(nb int) bool {
	if nb == 0 {
		return false
	}
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return false
		}
	}
	return true
}
