package main

//наш репозиторий, но нам нужна такая же реализация,
//как у суперрепозитория, потому что мы делаем с данными то же самое
type Repository interface {
	Save()
	Get()
	Update()
	Delete()
}
