package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	h := l.Head
	var data *interface{}
	for h != nil {
		if comp(ref, h.Data) {
			data = &h.Data
		}
		h = h.Next
	}
	return data
}
