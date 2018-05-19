package main

import (
	"time"
	"fmt"
)

func main()  {

	chUnbuffered := make(chan int)
	chbuffered := make(chan int, 100)

	go func() {
		for i := 1; i <= 10; i++ {
			// assign data
			chUnbuffered <- i
			chbuffered <- i
		}
	}()

	go func() {
		for {
			// exit data
			fmt.Println("chUnbuffered ===>", <- chUnbuffered)
			fmt.Println("chbuffered ===>", <- chUnbuffered)
		}
	}()
	time.Sleep(3 * time.Second)
}
