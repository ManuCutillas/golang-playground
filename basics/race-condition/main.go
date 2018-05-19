
package main

import (
"fmt"
"sync"
)

var wg sync.WaitGroup
var cycleCount int

func main() {
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go internCycle()
	}
	fmt.Println("The final countdown :", cycleCount)
}

func internCycle() {
	for i := 1; i <= 10; i++ {
		x := cycleCount
		x++
		cycleCount = x
	}
	wg.Done()
}
