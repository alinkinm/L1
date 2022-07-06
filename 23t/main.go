package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	//хотим удалить второй элемент
	i := 2

	//в слайс {cde} копируем {de}, т.е.: теперь вместо части {cde} будет стоять {de} - как бы накладываем второй аргумент
	//на первыйб т.к. два элемента накладываем на три, третий остается
	copy(a[i:], a[i+1:])

	//и его нужно удалить, т.к. это очевидно повторяющийся элемент
	a = a[:len(a)-1]

	fmt.Println(a)
}