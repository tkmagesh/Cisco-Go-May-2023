package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func main() {
	p := Product{100, "Pen", 10}
	fmt.Println(p)
}
