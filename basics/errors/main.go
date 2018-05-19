package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrBadStartUp = errors.New("Failed to start correctly")
)

type _error interface {
	Error() string
}

type MyError struct {
	str string
	when time.Time
}

func main() {
	standardErr()
	customError()
}

func standardErr() {
	var err _error
	err = ErrBadStartUp
	if err != ErrBadStartUp{
		fmt.Printf("Unknown Error -> %s\n", err)
		return
	}
	fmt.Printf("Error %s\n", err)
}

func (e MyError) Error() string {
	return fmt.Sprintf("%s at %s", e.str, e.when)
}

func customError(){
	err := MyError{}
	err.str = "Ohh No!"
	err.when = time.Now()
	printErr := err.Error()
	fmt.Println(printErr)
}
