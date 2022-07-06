package main

//должен наследоваться от суперрепозитория и имплементировать репозиторий
type Adapter struct {
	//наследуемся от суперрепозитория
	superRepo SuperRepository
}

//передаем в качестве параметра объект адаптера значит реализуем интерфейс репозитория
//а внутри вызываем соответствующий метод суперрепозитрия, от которого наследуемся
func Save(a Adapter) {
	SaveData(a.superRepo)
}

func Get(a Adapter) {
	GetData(a.superRepo)
}

func Update(a Adapter) {
	UpdateData(a.superRepo)
}

func Delete(a Adapter) {
	DeleteData(a.superRepo)
}
