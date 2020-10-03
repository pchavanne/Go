/*
	Commentaire multiligne
*/
package main

import (
	"fmt"
	"time"
)

func run(name string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name, ":", i)
	}
}

// commentaire simple ligne
func main() {
	start := time.Now()
	go run("Pop")
	go run("Jul")
	run("Max")
	end := time.Now()
	fmt.Println(end.Sub(start))
}
