package main

import (
	"fmt"
	"sync"
)

type counter struct {
	i int
}

func main() {
	//присваиваем начальное значение
	count := counter{i: 0}
	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	//допустим 10
	for k := 0; k < 10; k++ {
		wg.Add(1)
		go func(k int) {
			//кроме действующей горутины никто не может сейчас обратиться к переменной и поменять ее пока стоит лок
			m.Lock()
			count.i = k
			fmt.Println(count.i)
			//переменную поменяли, можно дать возможность остальным процессам продолжить
			m.Unlock()
			wg.Done()
		}(k)
	}
	wg.Wait()
}
