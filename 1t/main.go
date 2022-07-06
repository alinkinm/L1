package main

import "fmt"

type Human struct {
	name      string
	education string
}

type Action struct {
	Human
	action string
}

func (human Human) introduce() {
	fmt.Println("name: ", human.name, "\neducation:", human.education)
}

func (action Action) intoduceAction() {
	fmt.Println("action:", action.action)
}

//реализация встраивания методов
func (action Action) introduceExecutive() {
	fmt.Println("executive:")
	//у объекта структуры Action я могу вызвать метод структуры Human, т.е. это встроенный метод
	action.introduce()
}

func main() {
	action := Action{
		Human{
			"valeriy",
			"lawyer",
		},
		"teaching law",
	}

	action.intoduceAction()
	action.introduceExecutive()

}
