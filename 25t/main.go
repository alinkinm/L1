package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(6)
	sleep(5 * time.Second)
	fmt.Println(7)
}

func sleep(s time.Duration) {
	//засекаем время с момента вызова функции и по циклу каждые сколько-то наносекунд (наверное) меняем переменную
	// отвечающую за прошедшее время до тех пор, пока она не приблизится к нужной нам
	start := time.Now()
	since := time.Since(start)
	//не знаю временные промежутки в которые выполняется итерация, поэтому ровно в дурейшн не попадаю, поэтому
	//не !=, а <
	for since < s {
		since = time.Since(start)
	}
	return
}
