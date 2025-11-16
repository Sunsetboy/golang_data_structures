package main

import (
	"data_structures/structures"
	"fmt"
)

func main() {
	list := structures.NewList(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	fmt.Printf("List: %v \n", list.List())
	fmt.Printf("List joined: %s\n", list.Join())

	// create a binary tree from a slice

	binaryTreeElements := []int{1, 2, 3, 4, 5, 6}
	btree, err := structures.CreateBtreeFromSlice(binaryTreeElements)
	if err != nil {
		fmt.Printf("Error creating a binary tree: %s", err.Error())
	}
	fmt.Printf("Binary tree: %v", *btree)
}
