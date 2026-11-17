package main

import (
	"fmt"

	"piscine"
)

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	piscine.ForEach(piscine.PrintNbr, a)
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	result := piscine.Map(piscine.IsPrime, a)
// 	fmt.Println(result)
// }

// func main() {
// 	a1 := []string{"Hello", "how", "are", "you"}
// 	a2 := []string{"This", "is", "4", "you"}
// 	result1 := piscine.Any(piscine.IsNumeric, a1)
// 	result2 := piscine.Any(piscine.IsNumeric, a2)
// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }

// func main() {
// 	tab1 := []string{"Hello", "how", "are", "you"}
// 	tab2 := []string{"This", "1", "is", "4", "you"}
// 	answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
// 	answer2 := piscine.CountIf(piscine.IsNumeric, tab2)
// 	fmt.Println(answer1)
// 	fmt.Println(answer2)
// }

// func main() {
// 	a1 := []int{0, 1, 2, 3, 4, 5}
// 	a2 := []int{0, 2, 1, 3}
// 	result1 := IsSorted(f, a1)
// 	result2 := IsSorted(f, a2)
// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }

// func main() {
// 	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	piscine.SortWordArr(result)
// 	fmt.Println(result)
// }

// func main() {
// 	steps := piscine.CollatzCountdown(12)
// 	fmt.Println(steps)
// }

// func main() {
// 	middle := piscine.Abort(2, 3, 8, 5, 7)
// 	fmt.Println(middle)
// }

// func main() {
// 	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
// 	for index, element := range piscine.ShoppingSummaryCounter(summary) {
// 		fmt.Println(index, "=>", element)
// 	}
// }

// func main() {
// 	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// 	piscine.DealAPackOfCards(deck)
// }

// func main() {
// 	fmt.Print(piscine.JumpOver("1010101010"))
// 	fmt.Print(piscine.JumpOver(""))
// 	fmt.Print(piscine.JumpOver("t w e l v e"))
// 	fmt.Print(piscine.JumpOver("12"))
// }

// func main() {
// 	x := 5
// 	y := &x
// 	z := &y
// 	a := &z
// 	w := 2
// 	b := &w
// 	u := 7
// 	e := &u
// 	f := &e
// 	g := &f
// 	h := &g
// 	i := &h
// 	j := &i
// 	c := &j
// 	k := 6
// 	l := &k
// 	m := &l
// 	n := &m
// 	d := &n
// 	fmt.Println(***a)
// 	fmt.Println(*b)
// 	fmt.Println(*******c)
// 	fmt.Println(****d)
// 	piscine.Enigma(a, b, c, d)
// 	fmt.Println("After using Enigma")
// 	fmt.Println(***a)
// 	fmt.Println(*b)
// 	fmt.Println(*******c)
// 	fmt.Println(****d)
// }

// func main() {
// 	fmt.Println(piscine.DescendAppendRange(10, 5))
// 	fmt.Println(piscine.DescendAppendRange(5, 10))
// }

// func main() {
// 	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
// 	fmt.Println(piscine.ShoppingListSort(slice))
// }

func main() {
	fmt.Println(piscine.ActiveBits(7))
}
