package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type PersihableProduct struct {
	Product
	Expiry string
}

/*
	Create a "toString()" function that will return a formatted string of the given product
		'Id = 100, Name = "Apple", Cost = 50'

	Create a "applyDiscount" function that will apply the given discount on the given product
		applyDiscount(?, 10)
		product.Cost === 45

	Create the above functions for PerishableProduct as well

*/

//For Product
func (p Product) toString() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) applyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

//For PerishableProduct
func (pp PersihableProduct) toString() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.toString(), pp.Expiry)
}

func main() {
	// Product
	pen := &Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(toString(pen))
	fmt.Println(pen.toString())
	fmt.Println("After applying discount")
	// applyDiscount(&pen, 10)
	// (&pen).applyDiscount(10)
	pen.applyDiscount(10)

	// fmt.Println(toString(pen))
	fmt.Println(pen.toString())

	//PerishableProduct
	grapes := PersihableProduct{
		Product: Product{
			Id:   102,
			Name: "Grapes",
			Cost: 40,
		},
		Expiry: "2 Days",
	}
	fmt.Println(grapes.toString())
	fmt.Println("After applying discount")
	// applyDiscount(&(grapes.Product), 10)
	grapes.applyDiscount(10) // applyDiscount inherited from the Product (composed in PerishableProduct)
	fmt.Println(grapes.toString())
}
