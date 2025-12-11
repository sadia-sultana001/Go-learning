package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 5

// func add(a, b int) {
// fmt.Println(a + b)
// }
func print(num int) {
	fmt.Println("hello word", num)
	//add(3, 5)
}

func main() {
	var b = 13
	fmt.Println(b)
	go print(1)

	go print(2)

	go print(3)

	go print(4)

	go print(5)

	fmt.Println(a, " ", p)

	time.Sleep(5 * time.Second)

}
