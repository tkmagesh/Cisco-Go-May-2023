package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float32
}

func main() {
	/*
		var product struct {
			id   int
			name string
			cost float32
		} = struct {
			id   int
			name string
			cost float32
		}{
			id:   100,
			name: "Pen",
			cost: 10,
		}
	*/

	/*
		type Product struct {
			id   int
			name string
			cost float32
		}
	*/

	/*
		var product Product = Product{
			id:   100,
			name: "Pen",
			cost: 10,
		}
	*/

	// var product Product
	// var product Product = Product{100, "Pen", 10} // not advisable
	var product Product = Product{
		id:   100,
		name: "Pen",
		cost: 10,
	}
	fmt.Println(product)

	fmt.Printf("Id = %d, Name = %q, Cost = %.2f\n", product.id, product.name, product.cost)

	/*
		var p1 = Product{200, "Pencil", 5}
		var p2 = Product{200, "Pencil", 5}
		fmt.Println(p1 == p2) //=> EVERYTHING is a VALUE
	*/

	var productPtr *Product
	productPtr = &product
	// fmt.Printf("Id = %d, Name = %q, Cost = %.2f\n", (*productPtr).id, (*productPtr).name, (*productPtr).cost)
	fmt.Printf("Id = %d, Name = %q, Cost = %.2f\n", productPtr.id, productPtr.name, productPtr.cost) // NO NEED to dereference a pointer to a struct to access the members
}
