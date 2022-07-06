package main

import (
	"fmt"
	"sync"
)

func main() {
	x := [5]int8{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	n := 5
	//добавляем н горутин в вейтгруп для возможности ожидания
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			fmt.Println(x[i] * x[i])
			wg.Done()
		}(i)
	}

	//ждем пока все добавленные в вейтгруп горутины завершат работу
	wg.Wait()
}
