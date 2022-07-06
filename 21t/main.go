package main

func main() {
	superRepo := SuperRepository{}
	repoAdapter := Adapter{superRepo: superRepo}
	//получается вызвать методы у репоадаптера
	Save(repoAdapter)
	Get(repoAdapter)
	Update(repoAdapter)
	Delete(repoAdapter)

	//вопрос - мы все равно создали объект суперрепозитория, как объяснить смысл использования паттерна в этом
	//случае
}
