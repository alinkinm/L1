package main

import (
	"fmt"
	"sync"
	"time"
)

//race conditions +
func withoutChannels() {
	start := time.Now()
	x := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	n := 5
	wg.Add(n)
	sum := 0

	for i := 0; i < n; i++ {
		go func(i int) {
			sum = sum + x[i]
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println(sum)
}

func usingChannels() {
	x := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	n := 5
	nums := make(chan int)
	var sum int
	wg.Add(5)

	for i := 0; i < n; i++ {
		//параметр в функции для того, чтобы у каждой горутины был свой элемент массива для обработки
		go func(i int) {
			nums <- x[i] * x[i]
			wg.Done()
		}(i)
	}
	//никто не закрывает канал, эта горутина, т.к. она не занесена в WG просто перестает выполняться, как только
	//остальные горутины выполнились, поэтому не возникает ошибки, если мы слушаем пустой канал, потому что
	//программа просто завершается
	go func() {
		for num := range nums {
			sum = sum + num
			fmt.Println(sum)
		}
	}()

	wg.Wait()

}

func main() {
	usingChannels()
}
