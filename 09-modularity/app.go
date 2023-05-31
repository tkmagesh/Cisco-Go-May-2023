package main

import (
	"fmt"
	// "github.com/tkmagesh/cisco-go-may-2023/09-modularity/calculator"

	calc "github.com/tkmagesh/cisco-go-may-2023/09-modularity/calculator"
	"github.com/tkmagesh/cisco-go-may-2023/09-modularity/calculator/utils"

	"github.com/fatih/color"
	"github.com/tkmagesh/cisco-go-may-2023/09-modularity/models"
)

/* Attempting to create a method for the Product type in the "models" package */
// A method in the non-local type
/*
func (p models.Product) WhoAmI() string {
	return "I am a product"
}
*/

type MyProduct models.Product

func (p MyProduct) WhoAmI() string {
	return "I am a product"
}

func main() {
	color.Red("app invoked")

	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Op Count =", calc.OpCount())

	fmt.Println("IsEven(21) =", utils.IsEven(21))

	var p MyProduct
	fmt.Println(p.WhoAmI())
}
