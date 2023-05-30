package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	cisco := &Organization{
		Name: "Cisco",
		City: "Bangalore",
	}
	e1 := Employee{
		Id:   100,
		Name: "Magesh",
		Org:  cisco,
	}
	e2 := Employee{
		Id:   200,
		Name: "Suresh",
		Org:  cisco,
	}
	fmt.Println(e1)
	fmt.Println(e2)

	cisco.City = "Chennai"
	fmt.Println(e1.Org.City)
	fmt.Println(e2.Org.City)
}
