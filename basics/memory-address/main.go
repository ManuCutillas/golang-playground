package main

import "fmt"

func main() {
	getMemory()
}

func getMemory() {

	a := 20

	fmt.Println(a)
	fmt.Println(&a)

	a = 10

	fmt.Println(a)
	fmt.Println(&a)
}