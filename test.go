/*
	Commentaire multiligne
*/
package main

import (
	"fmt"
	"math"
	"time"
)

var i int = 1

func run(name string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name, ":", i)
	}
}

//Point struct
type Point struct {
	X float64
	Y float64
}

// Distance function
func Distance(p, q Point) float64 {
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}

// Distance method
func (p Point) Distance(q Point) float64 {
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}

// commentaire simple ligne
func main() {
	start := time.Now()
	// go run("Pop")
	// go run("Jul")
	// run("Max")
	p := Point{0, 0}
	q := Point{2, 2}
	d := Distance(p, q)
	fmt.Println("Distance: ", d)
	fmt.Println("Distance: ", p.Distance(q))
	end := time.Since(start).Milliseconds()
	fmt.Println("Took: ", end)
}
