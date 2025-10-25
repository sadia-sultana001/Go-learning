package main

import "fmt"

func main() {
	//pointer or address of memory
	x := 20

	y := 56

	p := &x
	q := &y

	*p = 40

	fmt.Println(p)
	fmt.Println(q)

	fmt.Println("X value: ", x)
	fmt.Println("Value od address: ", *p)
	fmt.Println("Value od address: ", *q)
}
