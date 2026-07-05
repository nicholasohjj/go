package main

import (
"fmt"
"math"
)

type Point struct {
	X,Y float64
}

func (p Point) Distance (q Point) float64 {
	return math.Sqrt(math.Pow(p.X-q.X,2) + math.Pow(p.Y-q.Y,2))
}

func (p *Point) Scale (s float64) {
	p.X *= s
	p.Y *= s
}

func main() {
	p,q := Point{1,2}, Point{3,4}
	fmt.Println(p.Distance(q))
	p.Scale(2)
	fmt.Println(p.Distance(q))
}
