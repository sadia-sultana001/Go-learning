package main

import "fmt"

func changeSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

func main() {

	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)
	a := x[4:]

	y := changeSlice(a)

	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(y)
	fmt.Println(x[0:8])

	/*	s := make([]int, 3, 5) //make function with len and cap. Systex p:= make([]int, 3,7)
		s[0] = 4
		s[2] = 19
		fmt.Println(s)
		fmt.Println("Slice", s, "Len", len(s), "Capaacity", cap(s))
		var x []int //empty slice []
		x = append(x, 1) //[1]
		x = append(x, 2)
		x = append(x, 3)
		y := x
		x = append(x, 4)
		y = append(y, 5)
		x[0] = 10
		fmt.Println(x)
		fmt.Println("Slice", x, "Len", len(x), "Capaacity", cap(x))
		fmt.Println(y)
	*/

}
