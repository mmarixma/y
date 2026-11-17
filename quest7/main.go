package main

import (
	"fmt"

	"piscine"
)

// func main() {
// 	fmt.Println(piscine.AppendRange(5, 10))
// 	fmt.Println(piscine.AppendRange(10, 5))
// }

// func main() {
// 	fmt.Println(piscine.MakeRange(5, 10))
// 	fmt.Println(piscine.MakeRange(10, 5))
// }

// func main() {
// 	test := []string{"Hello", "how", "are", "you?"}
// 	fmt.Println(piscine.ConcatParams(test))
// }

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
}

// func main() {
// 	a := piscine.SplitWhiteSpaces("Hello how are you?")
// 	piscine.PrintWordsTables(a)
// }
