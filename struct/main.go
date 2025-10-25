package main

import "fmt"

type user struct {
	name string //member variable or property

	age int
}

func main() {
	var user1 user

	user1 = user{
		name: "Alice",
		age:  21,
	}

	fmt.Println("Name:", user1.name)
	fmt.Println("Age:", user1.age)

	usser2 := user{ //instance of object
		name: "Bob",
		age:  25,
	}

	fmt.Println("Name:", usser2.name)
	fmt.Println("Age:", usser2.age)
}
