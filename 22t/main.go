package main

import "fmt"

func main() {
	fmt.Println(calculator(10, 5, "+"))
}

func calculator(a int32, b int32, o string) (res int32) {
	switch o {
	case "+":
		return a + b
	case "-":
		//по умолчанию вычтем второе из первого
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return
	}

}
