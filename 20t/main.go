//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func main() {
	originalS := "snow dog sun"
	//сплит по пробелам и запись в массив
	s := strings.Fields(originalS)
	//реверс элементов массива
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println(s)

}
