package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	x := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	ch1 := make(chan int)
	wg := &sync.WaitGroup{}
	wg1 := &sync.WaitGroup{}

	for _, y := range x {
		wg.Add(1)
		go func(y int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- y
			fmt.Println("отправил ", strconv.Itoa(y), "в первый канал")
		}(y, wg)
	}

	go func() {
		wg.Wait()
		//предупреждаем дедлок
		close(ch)
	}()

	for sq := range ch {
		wg1.Add(1)
		go func(sq int, wg1 *sync.WaitGroup) {
			defer wg1.Done()
			fmt.Println("получил ", strconv.Itoa(sq), "из первого канала, возвел в квадрат")
			mq := sq * sq
			ch1 <- mq
			fmt.Println("отправил ", strconv.Itoa(mq), "во второй канал")
		}(sq, wg1)
	}

	go func() {
		wg1.Wait()
		//предупреждаем дедлок
		close(ch1)
	}()

	for number := range ch1 {
		fmt.Println(number)
	}
}
