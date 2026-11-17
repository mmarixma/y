package main

import (
	//"fmt"

	"piscine"

	"github.com/01-edu/z01"
)

// func main() {
// 	z01.PrintRune(piscine.NRune("Hello!", 3))
// 	z01.PrintRune(piscine.NRune("Salut!", 2))
// 	z01.PrintRune(piscine.NRune("Bye!", -1))
// 	z01.PrintRune(piscine.NRune("Bye!", 5))
// 	z01.PrintRune(piscine.NRune("Ola!", 4))
// 	z01.PrintRune('\n')
// }

// func main() {
// 	s := "Hello 78 World!    4455 /"
// 	nb := piscine.AlphaCount(s)
// 	fmt.Println(nb)
// }

// func main() {
// 	fmt.Println(piscine.Index("Hello!", "l"))
// 	fmt.Println(piscine.Index("Salut!", "alu"))
// 	fmt.Println(piscine.Index("Ola!", "hOl"))
// }

// func main() {
// 	fmt.Println(piscine.Concat("Hello!", " How are you?"))
// }

// func main() {
// 	fmt.Println(piscine.IsUpper("HELLO"))
// 	fmt.Println(piscine.IsUpper("HELLO!"))
// }

// func main() {
// 	fmt.Println(piscine.IsLower("hello"))
// 	fmt.Println(piscine.IsLower("hello!"))
// }

// func main() {
// 	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
// 	fmt.Println(piscine.IsAlpha("HelloHowareyou"))
// 	fmt.Println(piscine.IsAlpha("What's this 4?"))
// 	fmt.Println(piscine.IsAlpha("Whatsthis4"))
// }

// func main() {
// 	fmt.Println(piscine.IsNumeric("010203"))
// 	fmt.Println(piscine.IsNumeric("01,02,03"))
// }

// func main() {
// 	fmt.Println(piscine.IsPrintable("Hello"))
// 	fmt.Println(piscine.IsPrintable("Hello\n"))
// }

// func main() {
// 	fmt.Println(piscine.ToUpper("Hello! How are you?"))
// }

// func main() {
// 	fmt.Println(piscine.ToLower("Hello! How are you?"))
// }

// func main() {
// 	piscine.PrintNbrInOrder(321)
// 	piscine.PrintNbrInOrder(0)
// 	piscine.PrintNbrInOrder(321)
// }

// func main() {
// 	fmt.Println(piscine.TrimAtoi("12345"))
// 	fmt.Println(piscine.TrimAtoi("str123ing45"))
// 	fmt.Println(piscine.TrimAtoi("012 345"))
// 	fmt.Println(piscine.TrimAtoi("Hello World!"))
// 	fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
// 	fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
// 	fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
// 	fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))
// }

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(piscine.Join(toConcat, ":"))
// }

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
