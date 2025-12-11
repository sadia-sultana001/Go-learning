package main

import "fmt"

func sum(a int, b int) (s int) {
	s = a + b
	return
}
func a() {
	i := 0

	fmt.Println("first", i)

	defer fmt.Println("second", i)

	i = i + 1

	fmt.Println("third", i)
	defer fmt.Println("fourth", i)

	return
}

func calculate() (result int) {
	fmt.Println("first", result)

	show := func() {
		result += 10
		fmt.Println("defer", result)
	}

	defer show()

	result = 5

	fmt.Println("second", result)
	return
}

func calc() int {

	result := 0

	fmt.Println("first", result)

	show := func() {
		result += 10
		fmt.Println("defer", result)
	}

	defer show()

	result = 5

	fmt.Println("second", result)
	return result
}

func main() {
	a()
	res := sum(3, 4)
	fmt.Println(res)
	c := calculate()
	fmt.Println("main first", c)

	d := calc()
	fmt.Println("main second", d)

}
