package main

import (
"fmt"
shared "./shared-library"
)

func main() {
 st := "lol"
 st = shared.ToUpperCase(st)
 fmt.Println(st)
}
