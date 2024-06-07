package main

func main() {

	var isTrue bool
	isTrue = false

	if isTrue {
		println("It's true!")
	} else {
		println("It's false!")
	}

	var name string
	name = "Bilbo"

	if name == "Bilbo" {
		println("Hello, Bilbo!")
	} else {
		println("Hello, stranger!")
	}

	name = "Frodo"

	// Switch
	switch name {
	case "Bilbo":
		println("Hello, Bilbo!")
	case "Frodo":
		println("Hello, Frodo!")
	default:
		println("Hello, stranger!")
	}

}
