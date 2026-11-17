package main

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 { // top left
				z01.PrintRune('o')
			} else if row == 0 && col == x-1 { // top right
				z01.PrintRune('o')
			} else if row == y-1 && col == 0 { // bottom left
				z01.PrintRune('o')
			} else if row == y-1 && col == x-1 { // bottom right
				z01.PrintRune('o')
			} else if row == 0 || row == y-1 { // horizontal
				z01.PrintRune('-')
			} else if col == 0 || col == x-1 { // vertical
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 { // top left
				z01.PrintRune('/')
			} else if row == 0 && col == x-1 { // top right
				z01.PrintRune('\\')
			} else if row == y-1 && col == 0 { // bottom left
				z01.PrintRune('\\')
			} else if row == y-1 && col == x-1 { // bottom right
				z01.PrintRune('/')
			} else if row == 0 || row == y-1 { // horizontal
				z01.PrintRune('*')
			} else if col == 0 || col == x-1 { // vertical
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 { // top left
				z01.PrintRune('A')
			} else if row == 0 && col == x-1 { // top right
				z01.PrintRune('A')
			} else if row == y-1 && col == 0 { // BOTTOM LEft
				z01.PrintRune('C')
			} else if row == y-1 && col == x-1 { // bottom right
				z01.PrintRune('C')
			} else if row == 0 || row == y-1 { // horizontz
				z01.PrintRune('B')
			} else if col == 0 || col == x-1 { // vertical
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')

	}
}

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				z01.PrintRune('A')
			} else if row == 0 && col == x-1 {
				z01.PrintRune('C')
			} else if row == y-1 && col == 0 {
				z01.PrintRune('A')
			} else if row == y-1 && col == x-1 {
				z01.PrintRune('C')
			} else if row == 0 || row == y-1 {
				z01.PrintRune('B')
			} else if col == 0 || col == x-1 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				z01.PrintRune('A')
			} else if row == 0 && col == x-1 {
				z01.PrintRune('C')
			} else if row == y-1 && col == 0 {
				z01.PrintRune('C')
			} else if row == y-1 && col == x-1 {
				z01.PrintRune('A')
			} else if row == 0 || row == y-1 {
				z01.PrintRune('B')
			} else if col == 0 || col == x-1 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadE(1, 5)
}
