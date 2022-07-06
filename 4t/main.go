package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	var n int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		go func(i int) {
			for data := range c {
				time.Sleep(time.Duration(i+1) * time.Second)
				a := <-c
				fmt.Println(data, " hello from worker ", i, "i'm reading every ", i+1, " seconds", a)
			}
		}(i)
	}

	for {
		rand.Seed(time.Now().UnixNano())
		c <- rand.Intn(100)
	}

	//как выбирать способ завершения работы всех воркеров..

}
