package binarytree

import "testing"

func TestBtInt(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Add(3)
	bt.Add(1)
	bt.Add(2)
	bt.Add(4)
	bt.Add(5)

	if bt.Root.Left.Value != 1 {
		t.Fatalf("expected 1, got %v", bt.Root.Left.Value)
	}

	if bt.Root.Left.Right.Value != 2 {
		t.Fatalf("expected 2, got %v", bt.Root.Left.Right.Value)
	}

	if bt.Root.Right.Value != 4 {
		t.Fatalf("expected 4, got %v", bt.Root.Right.Value)
	}

	if bt.Root.Right.Right.Value != 5 {
		t.Fatalf("expected 5, got %v", bt.Root.Right.Right.Value)
	}
}

func TestBtString(t *testing.T) {
	bt := NewBinaryTree[string]()
	bt.Add("c")
	bt.Add("a")
	bt.Add("b")
	bt.Add("d")
	bt.Add("e")

	if bt.Root.Left.Value != "a" {
		t.Fatalf("expected a, got %v", bt.Root.Left.Value)
	}

	if bt.Root.Left.Right.Value != "b" {
		t.Fatalf("expected b, got %v", bt.Root.Left.Right.Value)
	}

	if bt.Root.Right.Value != "d" {
		t.Fatalf("expected d, got %v", bt.Root.Right.Value)
	}

	if bt.Root.Right.Right.Value != "e" {
		t.Fatalf("expected e, got %v", bt.Root.Right.Right.Value)
	}
}
