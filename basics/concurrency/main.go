package main

import (
"fmt"
"sync"
"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2) // Agregamos el numero de go routines que usaremos.
	go primero()
	go segundo()
	wg.Wait()
}

func primero() {
	for i:= 1; i<= 50; i++ {
		fmt.Println(i, "Primero")
		time.Sleep(1 * time.Millisecond)
	}
	wg.Done()
}

func segundo() {
	for i := 1; i <= 50; i++ {
		fmt.Println(i, "Segundo")
		time.Sleep(2 * time.Millisecond)
	}
	wg.Done()
}