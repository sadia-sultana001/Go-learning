package main

import (
	"fmt"
)

var a = 30
var b = 20

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	var p int = 30
	var q int = 40
	add(p, q)
	add(a, b)
	add(a, p)
	if a > 0 {
		fmt.Println("a is greater than 0")
	}
	// Immediately invoke function expression (IIFE)
	func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}(a, q)
}
func init() {
	fmt.Println("init function called")
}
