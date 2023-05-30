package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Dummy struct {
	Name string
}
type PersihableProduct struct {
	// Id string //overriding the ID from the Product struct
	Product
	Dummy
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PersihableProduct {
	return &PersihableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {

	// grapes := PersihableProduct{Product{100, "Grapes", 50}, "2 Days"}
	grapes := PersihableProduct{
		// Id: "Arabian-Grapes",
		Product: Product{
			Id:   100,
			Name: "Grapes",
			Cost: 50,
		},
		Expiry: "2 Days",
	}
	fmt.Println(grapes)
	fmt.Println(grapes.Expiry)
	// fmt.Println(grapes.Product.Id)
	fmt.Println(grapes.Id)

	// fmt.Println(grapes.Name)
	fmt.Println(grapes.Product.Name)
	fmt.Println(grapes.Dummy.Name)

	apple := NewPerishableProduct(102, "Apple", 80, "1 week")
	fmt.Println(apple)
}
