package main

import (
	"fmt"
	"strings"
)

func checkOriginality(s string) (flag bool) {
	flag = true
	//тип символа показывал инт32
	symbols := map[int32]int{}
	//превращаем все в нижний регистр, чтобы инт32 у букв строки были одинаковые
	a := strings.ToLower(s)
	for _, ch := range a {
		//_, потому что нам не важно само значение, а ок возвращает булеан - есть ли уже такой ключ или нет
		//если он есть, остальное неважно
		if _, ok := symbols[ch]; ok {
			flag = false
		} else {
			symbols[ch] = 1
		}
		fmt.Println(ch)
	}
	return flag
}

func main() {
	fmt.Println(checkOriginality("abcdA"))

}
