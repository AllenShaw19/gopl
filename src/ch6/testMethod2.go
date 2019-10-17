package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) Distance(q *Point) float64 {
	return math.Sqrt(math.Pow(p.X-q.X, 2) + math.Pow(p.Y-q.Y, 2))
}
func PrintPoint(f func(*Point) float64, a *Point) float64 {

	fmt.Println("PrintPoint")
	return f(a)
}

func main() {

	p := &Point{
		1, 2,
	}

	q := &Point{
		3, 4,
	}

	X := 10
	Y := 20

	fmt.Println(p.Distance(q))

	f := p.Distance

	fmt.Println(f(q))

	fmt.Println(PrintPoint(f, q))

	fmt.Println(X + Y)

	day := 24 * time.Hour
	fmt.Printf("%T\n", day)
	//day.Seconds()
}
