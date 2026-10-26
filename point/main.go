package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int32
	y int32
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	arr := [15]rune{
		'x', ' ', '=', ' ',
		'0' + points.x/10, '0' + points.x%10, ' ',
		'y', ' ', '=', ' ',
		'0' + points.y/10, '0' + points.y%10, '\n',
	}

	for _, r := range arr {
		z01.PrintRune(r)
	}
}
