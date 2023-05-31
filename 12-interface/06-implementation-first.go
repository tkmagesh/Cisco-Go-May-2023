package main

import (
	"fmt"
	"math"
)

/* ---- Frozen ---- */
/* Implementation comes first */
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Length  float32
	Breadth float32
}

func (r Rectangle) Area() float32 {
	return r.Breadth * r.Length
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Breadth + r.Length)
}

/* ---- Frozen ---- */

/* contracts comes later */
type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

type PerimeterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimter :", x.Perimeter())
}

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintStats(x ShapeStatsFinder) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	r := Rectangle{10, 20}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)

}
