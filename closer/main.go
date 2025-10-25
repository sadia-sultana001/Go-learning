package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age:", age)

	show := func() {
		money = money + age + p + a //moneey -> escape analysis -> heapff
		fmt.Println(money)
	}
	return show
}

func call() {
	incrl := outer()
	incrl()
	incrl()

	incr2 := outer()
	incr2()
	incr2()
}

func init() {
	fmt.Println("Init function called")

}
func main() {
	call()
}
