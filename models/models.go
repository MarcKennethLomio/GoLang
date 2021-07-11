package models

type Person struct {
	Name string
}

func(person *Person) GetPersonName() string {
	return person.Name
}