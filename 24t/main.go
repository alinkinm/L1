package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func NewPoint(x int, y int) *point {
	return &point{
		x: x,
		y: y,
	}
}

func findDistance(a *point, b *point) (distance float64) {
	ac := b.x - a.x
	bc := b.y - a.y
	return math.Sqrt((float64(ac*ac + bc*bc)))
}

func main() {
	a := NewPoint(10, 25)
	b := NewPoint(5, 1)
	fmt.Println(findDistance(a, b))
}
