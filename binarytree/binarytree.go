package binarytree

import "cmp"

type BinaryTree[T cmp.Ordered] struct {
	Root *Node[T]
}

func NewBinaryTree[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (b *BinaryTree[T]) Add(value T) {
	if b.Root == nil {
		b.Root = &Node[T]{Value: value}
	} else {
		b.Root.add(value)
	}
}

type Node[T cmp.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (n *Node[T]) add(value T) {
	if value < n.Value {
		n.addToLeft(value)
	} else if n.Value < value {
		n.addToRight(value)
	}
}

func (n *Node[T]) addToLeft(value T) {
	if n.Left == nil {
		n.Left = &Node[T]{Value: value}
	} else {
		n.Left.add(value)
	}
}

func (n *Node[T]) addToRight(value T) {
	if n.Right == nil {
		n.Right = &Node[T]{Value: value}
	} else {
		n.Right.add(value)
	}
}
