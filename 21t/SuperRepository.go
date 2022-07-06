package main

import "fmt"

type SuperRepository struct {
}

func SaveData(repository SuperRepository) {
	fmt.Println("saving data to repository")
}

func GetData(repository SuperRepository) {
	fmt.Println("getting data from repository")
}

func UpdateData(repository SuperRepository) {
	fmt.Println("updating data in repository")
}

func DeleteData(repository SuperRepository) {
	fmt.Println("deleting data from repository")
}
