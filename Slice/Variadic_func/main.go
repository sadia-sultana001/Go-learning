package main

import "fmt"

// variadic function
func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
func main() {
	print(4, 5, 6, 8, 12, 11)
}
