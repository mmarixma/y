package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func getLength(head *NodeI) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

func ListSort(l *NodeI) *NodeI {
	// buublesort
	head := l
	size := getLength(l)
	itr := 0

	for itr < size {
		curr := head
		for j := 0; j < size-itr-1; j++ {
			if curr.Data > curr.Next.Data {
				data := curr.Data
				curr.Data = curr.Next.Data
				curr.Next.Data = data
			}
			curr = curr.Next
		}
		itr++
	}
	return head
}
