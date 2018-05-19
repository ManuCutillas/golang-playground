package main

import (
	"fmt"
)

func main()  {

	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			// assign data
			ch <- i
		}
		// close channel
		close(ch)
	}()

		for y := range ch {
			// exit data
			fmt.Println("channel value ===>", y)
		}
}
