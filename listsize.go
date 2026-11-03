package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListSize(l *List) int {
	count := 0
	t := l.Head
	for t != nil {
		count++
		t = t.Next
	}
	return count
}

// func ListPushFront(l *List, data interface{}) {
// 	n := &NodeL{Data: data}
// 	if l.Head == nil {
// 		l.Head = n
// 		l.Tail = n
// 		return
// 	}
// 	n.Next = l.Head
// 	l.Head = n
// }
