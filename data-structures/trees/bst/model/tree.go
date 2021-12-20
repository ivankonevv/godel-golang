package model

import "fmt"

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(data int) {
	if t.Root == nil {
		t.Root = &Node{Key: data}
		return
	}
	t.Root.Insert(data)
}

func (t *Tree) Find(key int) (*Node, error) {
	if t.Root == nil {
		return nil, fmt.Errorf("root not found. The tree is empty")
	}
	return t.Root.find(key)
}

func (t *Tree) Remove(key int) error {
	if t.Root == nil {
		return fmt.Errorf("root not found. The tree is empty")
	}
	return remove(&t.Root, key)
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(data int) {
	if data < n.Key {
		// insert into the left tree
		if n.Left == nil {
			n.Left = &Node{Key: data}
		} else {
			n.Left.Insert(data)
		}
		return
	}
	// insert into the right tree
	if n.Right == nil {
		n.Right = &Node{Key: data}
	} else {
		n.Right.Insert(data)
	}
}

func (n *Node) find(key int) (*Node, error) {
	if n == nil {
		return nil, fmt.Errorf("node does not exist")
	}
	if n.Key > key {
		return n.Left.find(key)
	}
	if n.Key < key {
		return n.Right.find(key)
	}
	return n, nil
}

func remove(node **Node, key int) error {
	n := *node
	switch {
	case n == nil:
		return fmt.Errorf("node does not exist")
	case n.Key > key:
		return remove(&n.Left, key)
	case n.Key < key:
		return remove(&n.Right, key)
	}
	if n.Left != nil && n.Right != nil {
		minParent := n
		minRight := n.Right

		for minRight.Left != nil {
			minParent = minRight
			minRight = minRight.Left
		}

		if minParent != n {
			minParent.Left = minRight.Right
		} else {
			minParent.Right = minRight.Right
		}

		tmp := *minRight
		tmp.Left = n.Left
		tmp.Right = n.Right
		*node = &tmp
		return nil
	}
	if n.Left == nil && n.Right == nil {
		*node = nil
		return nil
	}
	if n.Left != nil {
		*node = (*node).Left
	} else {
		*node = (*node).Right
	}
	return nil
}
