package main

import (
	"fmt"
	// "github.com/tkmagesh/cisco-go-may-2023/09-modularity/calculator"

	calc "github.com/tkmagesh/cisco-go-may-2023/09-modularity/calculator"
	"github.com/tkmagesh/cisco-go-may-2023/09-modularity/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	color.Red("app invoked")

	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Op Count =", calc.OpCount())

	fmt.Println("IsEven(21) =", utils.IsEven(21))
}
