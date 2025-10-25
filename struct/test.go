package main

import "fmt"

type user struct {
	name string //member variable or property

	age int
}

/*func printUserDetails(u user) {
	fmt.Println("Name:", u.name)
	fmt.Println("Age:", u.age)
}
*/
// receiver function
func (u user) printDetails() {
	fmt.Println("Name: ", u.name)
	fmt.Println("Age: ", u.age)
}

func (u user) call(a int) {
	fmt.Println(u.name)
	fmt.Println(a)
}

func main() {
	var user1 user

	user1 = user{
		name: "Alice",
		age:  21,
	}

	//	printUserDetails(user1)

	user1.printDetails()
	user1.call(10)

	usser2 := user{ //instance of object
		name: "Bob",
		age:  25,
	}

	//	printUserDetails(usser2)
	usser2.printDetails()
	usser2.call(20)
}
