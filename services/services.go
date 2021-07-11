package services

import (
	models "models"
)

func GetPersonName(){
	person := models.Person{
		Name: "Marc Kenneth Lomio",
	}

	println(person.GetPersonName())
}