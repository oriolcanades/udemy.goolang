package main

import "fmt"

type Animal interface {
	Speak() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Fido",
		Breed: "Golden Retriever",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 32,
	}

	PrintInfo(&gorilla)

}

func PrintInfo(animal Animal) {
	fmt.Println("This animal says", animal.Speak(), "and has", animal.NumberOfLegs(), "legs")
}

func (d *Dog) Speak() string {
	return "Woof!"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Speak() string {
	return "Ugh!"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
