package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("Labore exercitation nostrud ex minim et duis consequat. Magna elit dolore laboris excepteur. Enim ad est magna sit cillum dolore deserunt adipisicing magna officia labore quis exercitation cillum. Officia laborum eiusmod minim et veniam qui adipisicing sunt veniam eu fugiat voluptate ex.")
	fmt.Println(str.Length())
}
