package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	if l == nil {
		return newNode
	}

	head := l
	start := l
	var prev *NodeI

	for head != nil {
		if data_ref < head.Data {
			if head == l {
				newNode.Next = head
				return newNode
			} else {
				prev.Next = newNode
				newNode.Next = head
				return start
			}
		}
		prev = head
		head = head.Next
	}
	prev.Next = newNode
	return start
}
