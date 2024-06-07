package main

import (
	//"fmt" // pronounced 'format
	"log"
	"sort"
	"time"
)

//variable package level
//var name string = "Bilbo"

func main() {
	//fmt.Println("Hello, World!")

	// Variables
	//var name string = "Bilbo"
	//var number int = 42

	bilboName := getName()
	street, number := getStreetAndNumber()
	log.Println("Hello,", bilboName, ". Your address is:", street, number)

	// Pointers
	log.Println("Before:", bilboName)
	chargeUsingPointer(&bilboName)
	log.Println("Changed to:", bilboName)

	//
	log.Println(setName("Frodo"))

	person := Person{
		Name:      "Frodo",
		Age:       33,
		Email:     "frondo@example.com",
		BirthDate: time.Date(1990, time.January, 15, 8, 34, 59, 3, time.UTC),
	}

	log.Println(person.Email)
	log.Println(person.getName())

	person.printPerson()

	// Maps ---------- key--- value
	// Maps are not ordered
	myMap := make(map[string]string)
	myMap["name"] = "Bilbo"
	log.Println(myMap["name"])

	personMap := make(map[string]Person)
	personMap["Frodo"] = Person{
		Name:      "Frodo",
		Age:       33,
		Email:     "frodo@baggins.com",
		BirthDate: time.Date(1990, time.January, 15, 8, 34, 59, 3, time.UTC),
	}
	log.Println(personMap["Frodo"].Email)

	// Slices
	// Slices are ordered
	// Slices are like arrays, but they can grow and shrink
	mySlice := []string{"Bilbo", "Frodo", "Gandalf"}
	log.Println(mySlice[1])

	var mySlice2 []string
	mySlice2 = append(mySlice2, "Bilbo")
	mySlice2 = append(mySlice2, "Frodo")
	mySlice2 = append(mySlice2, "Gandalf")
	log.Println(mySlice2)

	var mySlice3 []int
	mySlice3 = append(mySlice3, 1)
	mySlice3 = append(mySlice3, 3)
	mySlice3 = append(mySlice3, 2)
	log.Println(mySlice3)
	sort.Ints(mySlice3)
	log.Println(mySlice3)
	log.Println(mySlice3[0:2])

}

func getName() string {
	return "Bilbo"
}

func getStreetAndNumber() (string, int) {
	return "Bag End", 42
}

func chargeUsingPointer(s *string) {
	// s is the memory address of the variable, for example 0xc0000b8000
	log.Println("Pointer:", s)
	newValue := "Gandalf"
	*s = newValue
}

func setName(s string) (string, string) {
	return s, "Baggins"
}

type Person struct {
	Name      string
	Age       int
	Email     string
	BirthDate time.Time
}

func (p *Person) getName() string {
	return p.Name
}

func (p *Person) printPerson() {
	log.Println(*p)
}
