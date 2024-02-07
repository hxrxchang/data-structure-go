package binarytreemap

import "fmt"

type BinaryTreeMap[T any] struct {
	Root *Node[T]
}

func NewBinaryTreeMap[T any]() *BinaryTreeMap[T] {
	return &BinaryTreeMap[T]{}
}

func (b *BinaryTreeMap[T]) Add(key string, value T) {
	if b.Root == nil {
		b.Root = &Node[T]{Key: key, Value: value}
	} else {
		b.Root.add(key, value)
	}
}

func (b *BinaryTreeMap[T]) Get(key string) (T, error) {
	if b.Root == nil {
		var zero T
		return zero, fmt.Errorf("key: %s not found", key)
	}
	return b.Root.get(key)
}

type Node[T any] struct {
	Key string
	Value T
	Left *Node[T]
	Right *Node[T]
}

func (n *Node[T]) add(key string, value T) {
	if key == n.Key {
		n.Value = value
	} else if key < n.Key {
		n.addToLeft(key, value)
	} else {
		n.addToRight(key, value)
	}
}

func (n *Node[T]) addToLeft(key string, value T) {
	if n.Left == nil {
		n.Left = &Node[T]{Key: key, Value: value}
	} else {
		n.Left.add(key, value)
	}
}

func (n *Node[T]) addToRight(key string, value T) {
	if n.Right == nil {
		n.Right = &Node[T]{Key: key, Value: value}
	} else {
		n.Right.add(key, value)
	}
}

func (n *Node[T]) get(key string) (T, error) {
	if key == n.Key {
		return n.Value, nil
	} else if key < n.Key {
		if n.Left == nil {
			var zero T
			return zero, fmt.Errorf("key: %s not found", key)
		}
		return n.Left.get(key)
	} else {
		if n.Right == nil {
			var zero T
			return zero, fmt.Errorf("key: %s not found", key)
		}
		return n.Right.get(key)
	}
}
