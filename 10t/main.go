package main

import (
	"fmt"
	"math"
)

func main() {
	x := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	min := x[0]
	for i, e := range x {
		if i == 0 || e < min {
			min = e
		}
	}
	max := x[0]
	for i, e := range x {
		if i == 0 || e > max {
			max = e
		}
	}
	//находим ближайший десяток к минимуму и максимуму
	var left int
	left = int(min)
	for left%10 != 0 {
		left -= 1
	}
	right := int(max)
	for right%10 != 0 {
		right += 1
	}
	diff := (right - left) / 10
	mp := map[int][]float64{}
	for i := 0; i < diff; i++ {
		right = left + 10
		ne := []float64{}
		for k := range x {

			if math.Abs(x[k]-float64(left)) < 10.0 && x[k] > float64(left) {
				ne = append(ne, x[k])
			}
		}
		mp[left] = ne
		left += 10
	}
	fmt.Println(mp)
}
