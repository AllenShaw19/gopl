package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

type point struct {
	X int
	Y int
}

type circle struct {
	point
	Radius int
}

type wheel struct {
	circle
	Spokes int
}

func main() {
	Test2()
	//TestUnnamedStruct()
}

func Test2() {
	var w wheel
	w.X = 1
	w.Y = 2
	w.Radius = 3
	w.Spokes = 4

	var w2 wheel
	w2.circle.point.X = 1
	w2.circle.point.Y = 2
	w2.circle.Radius = 3
	w2.Spokes = 4

	fmt.Printf("w: %#v\nw1: %#v\n", w, w2)
	fmt.Println("w == w1 ?", w == w2)
}

func TestUnnamedStruct() {
	var w = &Wheel{}
	w.X = 123
	w.Y = 321
	w.Radius = 12
	w.Spokes = 6743
	w1 := &Wheel{Circle{Point{123, 321}, 12}, 6743}
	fmt.Printf("w: %#v\nw1: %#v\n", w, w1)
	fmt.Println("结构体比较 *w == *w1 ?", *w == *w1)
	fmt.Println("指针比较 w == w1 ?", w == w1)
}
