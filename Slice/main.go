package main

import "fmt"

var arr2 = [5]string{"I", "am", "Sadia", "and", "you"}

func main() {

	s := arr2[1:4]       //existing array
	s2 := s[1:2]         //Slice from a slice
	s1 := []int{1, 2, 3} //slice literal

	p := make([]int, 4) //make function with len
	p[0] = 2
	p[3] = 10

	fmt.Println(arr2)
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println("Slice", s1, "Len", len(s), "Capaacity", cap(s1))

	fmt.Println(p)
	fmt.Println(len(p))
	fmt.Println(cap(p))

}
