package main

import (
	"data_structures/structures"
	"fmt"
)

func main() {
	list := structures.NewListInt(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	fmt.Printf("List: %v", list.List())
	fmt.Printf("List joined: %s", list.Join())
}
