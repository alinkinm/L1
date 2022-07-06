package main

import "fmt"

func main() {
	array1 := []int{30, -568, 1, 0, -3500, 28, 99}
	array2 := []int{25, 22, 10, 99, 76, -1123, 0}

	mp := map[int]int{}

	for i := 0; i < len(array1); i++ {
		if val, ok := mp[array1[i]]; ok {
			val++
			mp[array1[i]] = val
		} else {
			mp[array1[i]] = 1
		}
	}

	for i := 0; i < len(array2); i++ {
		if val, ok := mp[array2[i]]; ok {
			val++
			mp[array2[i]] = val
		} else {
			mp[array2[i]] = 1
		}
	}

	fmt.Println(mp)

}
