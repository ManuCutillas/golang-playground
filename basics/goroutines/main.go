package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go getOne()
	go getTwo()
	wg.Wait()
}

func getOne() {
	for i := 1; i <= 50; i++ {
		fmt.Println("This first :", i)
		time.Sleep(1 * time.Millisecond)
	}
	wg.Done()
}

func getTwo() {
	for s := 1; s <= 50; s++ {
		fmt.Println("This  second :", s)
		time.Sleep(5 * time.Millisecond)
	}
	wg.Done()
}