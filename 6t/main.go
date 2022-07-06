package main

import (
	"fmt"
	"time"
)

func main() {

	//останавливаем в мейне
	go func() {
		fmt.Println("hello")
	}()

	time.Sleep(10 * time.Second)
}

func stoppingFromOtherGoroutine() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("hello")
			}
		}
	}()
	//произвольное время до остановки
	time.Sleep(5 * time.Second)
	quit <- true
}
