package main

import "fmt"

//pass by value
//pass by reference

type User struct {
	Name    string
	Age     int
	Salary  float64
	Favfood []string
}

func print(num *[5]int) {
	fmt.Println(num)
}

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	print(&arr)

	user := User{
		Name:   "John",
		Age:    30,
		Salary: 50000.50,
	}
	p := &user
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.Salary)

}
