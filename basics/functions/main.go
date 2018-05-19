package main

import "fmt"

var globalScope = "This is the global scope"

func main() {
	xy := []int{3,5,10,20,15}
	sum(xy...)
	clousure()
	recursion(4)
	defered()
}

func sum(numbers ...int) int {
	var counter int
	for _, v := range numbers {
		counter += v
	}
	fmt.Println(counter)
	fmt.Println(globalScope)
	return counter
}

func clousure() {
	shared := 0
	addOne := func() int {
		shared++
		return shared
	}
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())
}

func recursion(n int) int {
	fmt.Println("n value ====>", n)
	if n == 0 {
		return 1
	}
	return n * recursion(n - 1)
}

func defered() {
	one := 10
	two := 10
	sumSync(one, two)
	defer rest(one, two)
	multi(one, two)
}

func sumSync(x, y int) {
fmt.Println(x + y)
}

func rest(x, y int) {
	fmt.Println(x - y)
}

func multi(x, y int) {
	fmt.Println(x * y)
}