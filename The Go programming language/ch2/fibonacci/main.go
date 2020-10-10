package main

import (
	"fmt"
	"time"
)

type toto int

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	start := time.Now()
	fmt.Println(fib(100))
	fmt.Println(time.Since(start))
	i := 1
	fmt.Printf("%v, %T", toto(i), toto(i))
}
