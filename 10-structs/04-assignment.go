package main

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
func main() {

}
