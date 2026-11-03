package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	for pos > 0 && l != nil {
		l = l.Next
		pos--
	}
	return l
}
