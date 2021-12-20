package main

import (
	"fmt"
	"github.com/apaliavy/godel-golang/data-structures/trees/bst/model"
)

func main() {
	var t model.Tree

	t.Insert(10)
	t.Insert(1)
	t.Insert(15)
	t.Insert(16)
	t.Insert(8)
	t.Insert(10)
	t.Insert(3)
	key := 8
	node, err := t.Find(key)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("finded node: %#v\n", node)

	fmt.Printf("Remove key: %d\n", key)
	fmt.Println("Before: ")
	printPreOrder(t.Root)

	if err := t.Remove(key); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println("\nAfter: ")
	printPreOrder(t.Root)
}

func printPreOrder(n *model.Node) {
	if n == nil {
		return
	}

	fmt.Printf("%d ", n.Key)
	printPreOrder(n.Left)
	printPreOrder(n.Right)
}
