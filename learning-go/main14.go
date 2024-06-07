package main

func main() {

	// loops
	for i := 0; i < 5; i++ {
		println(i)
	}

	// range over a slice
	nameList := []string{"Bilbo", "Frodo", "Sam", "Merry", "Pippin"}

	// i = index
	// name = value
	for i, name := range nameList {
		println(name, "-", i)
	}

	// don't care about the index
	// use _ to ignore it
	for _, name := range nameList {
		println(name)
	}

	animals := map[string]string{"bird": "eagle", "mammal": "elephant", "reptile": "snake"}
	for animalType, animal := range animals {
		println(animal, "is a", animalType)
	}

	var firstLine = "Once upon a mistime dreary"
	firstLine = "x"
	for i, l := range firstLine {
		println(i, ":", l)
	}

	var users []User
	users = append(users, User{FirstName: "Bilbo", LastName: "Baggins", Age: 111})
	users = append(users, User{FirstName: "Frodo", LastName: "Baggins", Age: 33})
	users = append(users, User{FirstName: "Gandalf", LastName: "The Grey", Age: 2019})
	for _, user := range users {
		println(user.FirstName, user.LastName, "-", user.Age)
	}

}

type User struct {
	FirstName string
	LastName  string
	Age       int
}
