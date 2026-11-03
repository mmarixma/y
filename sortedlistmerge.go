package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}

	if n2 == nil {
		return n1
	}

	curr_head := n1
	var prev *NodeI

	for n1 != nil && n2 != nil {
		if n1.Data > n2.Data {
			nextN2 := n2.Next
			if n1 == curr_head {
				curr_head = n2
				n2.Next = n1
				prev = n2
			} else {
				prev.Next = n2
				n2.Next = n1
				prev = n2
			}
			n2 = nextN2
		} else {
			prev = n1
			n1 = n1.Next
		}
	}

	if n1 == nil && n2 != nil {
		prev.Next = n2
	}
	return curr_head
}
