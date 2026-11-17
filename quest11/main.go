package main

import (
	"fmt"
	"piscine"
)

func main() {
	link := &piscine.List{}
	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, "man")
	piscine.ListPushBack(link, "how are you")
	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}

// func main() {
// 	link := &piscine.List{}
// 	piscine.ListPushFront(link, "Hello")
// 	piscine.ListPushFront(link, "man")
// 	piscine.ListPushFront(link, "how are you")
// 	it := link.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " ")
// 		it = it.Next
// 	}
// 	fmt.Println()
// }

// func main() {
// 	link := &piscine.List{}
// 	link2 := &piscine.List{}
// 	piscine.ListPushBack(link, "three")
// 	piscine.ListPushBack(link, 3)
// 	piscine.ListPushBack(link, "1")
// 	fmt.Println(piscine.ListLast(link))
// 	fmt.Println(piscine.ListLast(link2))
// }

// type List = piscine.List
// type Node = piscine.NodeL
// func PrintList(l *List) {
// 	link := l.Head
// 	for link != nil {
// 		fmt.Print(link.Data, " -> ")
// 		link = link.Next
// 	}
// 	fmt.Println(nil)
// }
// func main() {
// 	link := &List{}
// 	piscine.ListPushBack(link, "I")
// 	piscine.ListPushBack(link, 1)
// 	piscine.ListPushBack(link, "something")
// 	piscine.ListPushBack(link, 2)
// 	fmt.Println("------list------")
// 	PrintList(link)
// 	piscine.ListClear(link)
// 	fmt.Println("------updated list------")
// 	PrintList(link)
// }

// func main() {
// 	link := &piscine.List{}
// 	piscine.ListPushBack(link, "hello")
// 	piscine.ListPushBack(link, "how are")
// 	piscine.ListPushBack(link, "you")
// 	piscine.ListPushBack(link, 1)
// 	fmt.Println(piscine.ListAt(link.Head, 3).Data)
// 	fmt.Println(piscine.ListAt(link.Head, 1).Data)
// 	fmt.Println(piscine.ListAt(link.Head, 7))
// }

// func main() {
// 	link := &piscine.List{}
// 	piscine.ListPushBack(link, 1)
// 	piscine.ListPushBack(link, 2)
// 	piscine.ListPushBack(link, 3)
// 	piscine.ListPushBack(link, 4)
// 	piscine.ListReverse(link)
// 	it := link.Head
// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}
// 	fmt.Println("Tail", link.Tail)
// 	fmt.Println("Head", link.Head)
// }

// func main() {
// 	link := &piscine.List{}
// 	piscine.ListPushBack(link, "1")
// 	piscine.ListPushBack(link, "2")
// 	piscine.ListPushBack(link, "3")
// 	piscine.ListPushBack(link, "5")
// 	piscine.ListForEach(link, piscine.Subtract3_node)
// 	piscine.ListForEach(link, piscine.Add2_node)
// 	piscine.ListForEach(link, piscine.Subtract3_node)
// 	piscine.ListForEach(link, piscine.Add2_node)
// 	piscine.ListForEach(link, piscine.Subtract3_node)
// 	piscine.ListForEach(link, piscine.Add2_node)
// 	it := link.Head
// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}
// }

// func PrintElem(node *piscine.NodeL) {
// 	fmt.Println(node.Data)
// }
// func StringToInt(node *piscine.NodeL) {
// 	node.Data = 2
// }
// func PrintList(l *piscine.List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, "->")
// 		it = it.Next
// 	}
// 	fmt.Print("nil", "\n")
// }
// func main() {
// 	link := &piscine.List{}
// 	piscine.ListPushBack(link, 1)
// 	piscine.ListPushBack(link, "hello")
// 	piscine.ListPushBack(link, 3)
// 	piscine.ListPushBack(link, "there")
// 	piscine.ListPushBack(link, 23)
// 	piscine.ListPushBack(link, "!")
// 	piscine.ListPushBack(link, 54)
// 	PrintList(link)
// 	fmt.Println("--------function applied--------")
// 	piscine.ListForEachIf(link, PrintElem, piscine.IsPositiveNode)
// 	piscine.ListForEachIf(link, StringToInt, piscine.IsAlNode)
// 	fmt.Println("--------function applied--------")
// 	PrintList(link)
// 	fmt.Println()
// }

// func main() {
// 	link := &piscine.List{}
// 	piscine.ListPushBack(link, "hello")
// 	piscine.ListPushBack(link, "hello1")
// 	piscine.ListPushBack(link, "hello2")
// 	piscine.ListPushBack(link, "hello3")
// 	found := piscine.ListFind(link, interface{}("hello2"), piscine.CompStr)
// 	fmt.Println(found)
// 	fmt.Println(*found)
// }

// func PrintList(l *piscine.List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }
// func main() {
// 	link := &piscine.List{}
// 	link2 := &piscine.List{}
// 	piscine.ListPushBack(link, "a")
// 	piscine.ListPushBack(link, "b")
// 	piscine.ListPushBack(link, "c")
// 	piscine.ListPushBack(link, "d")
// 	fmt.Println("-----first List------")
// 	PrintList(link)
// 	piscine.ListPushBack(link2, "e")
// 	piscine.ListPushBack(link2, "f")
// 	piscine.ListPushBack(link2, "g")
// 	piscine.ListPushBack(link2, "h")
// 	fmt.Println("-----second List------")
// 	PrintList(link2)
// 	fmt.Println("-----Merged List-----")
// 	piscine.ListMerge(link, link2)
// 	PrintList(link)
// }


//  func PrintList(l *piscine.NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }
// func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
// 	n := &piscine.NodeI{Data: data}
// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }
// func main() {
// 	var link *piscine.NodeI
// 	link = listPushBack(link, 5)
// 	link = listPushBack(link, 4)
// 	link = listPushBack(link, 3)
// 	link = listPushBack(link, 2)
// 	link = listPushBack(link, 1)
// 	PrintList(piscine.ListSort(link))
// }


// func PrintList(l *piscine.NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }
// func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
// 	n := &piscine.NodeI{Data: data}
// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }
// func main() {
// 	var link *piscine.NodeI
// 	link = listPushBack(link, 1)
// 	link = listPushBack(link, 4)
// 	link = listPushBack(link, 9)
// 	PrintList(link)
// 	link = piscine.SortListInsert(link, -2)
// 	link = piscine.SortListInsert(link, 2)
// 	PrintList(link)
// }

// func PrintList(l *piscine.NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }
// func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
// 	n := &piscine.NodeI{Data: data}
// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }
// func main() {
// 	var link *piscine.NodeI
// 	var link2 *piscine.NodeI
// 	link = listPushBack(link, 3)
// 	link = listPushBack(link, 5)
// 	link = listPushBack(link, 7)
// 	link2 = listPushBack(link2, -2)
// 	link2 = listPushBack(link2, 9)
// 	PrintList(piscine.SortedListMerge(link2, link))
// }


