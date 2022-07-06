package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("До сортировки: \n", slice)
	quicksort(slice)
	fmt.Println("Отсортированный: \n", slice)
}

//случайный массив
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	//ставим pivot на место последнего элемента
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		//проходимся по массиву и все меньше числа отправляем налево. лефт - стена
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			//отметка - слева будут находиться цифры меньше, справа - больше
			//инкрементируем каждый раз когда попадается число меньше пивота и ставится в этот индекс
			left++
		}
	}

	//ставим пивот на индекс лефт,т.к. изначально сравнивали с ним и теперь он на своем месте
	a[left], a[right] = a[right], a[left]

	//пивот оставляем на месте, и делаем сортировку для левой и правой частей массива
	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
