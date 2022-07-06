package main

import "fmt"

func main() {

	//как я поняла, создать мэп из последовательность, где в качестве ключей ее элементы, а значений - количество их повторений в ней
	catsdogs := []string{"cat", "cat", "dog", "cat", "tree"}
	newcats := map[string]int{}
	for x := range catsdogs {
		if val, ok := newcats[catsdogs[x]]; ok {
			val++
			newcats[catsdogs[x]] = val
		} else {
			newcats[catsdogs[x]] = 1
		}
	}

	fmt.Println(newcats)
}
