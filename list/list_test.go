package list

import (
	"testing"
)

func TestUnshift(t *testing.T) {
	l := NewList[int]()
	l.Unshift(4)
	l.Unshift(3)
	l.Unshift(1)
	l.Unshift(2)

	expected := []int{2, 1, 3, 4}
	got := l.Each()

	if len(got) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}

	for i := 0; i < len(expected); i++ {
		if got[i].Value != expected[i] {
			t.Fatalf("expected %v, got %v", expected[i], got[i].Value)
		}
	}
}

func TestPush(t *testing.T) {
	l := NewList[int]()
	l.Push(4)
	l.Push(3)
	l.Push(1)
	l.Push(2)

	expected := []int{4, 3, 1, 2}
	got := l.Each()

	if len(got) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}

	for i := 0; i < len(expected); i++ {
		if got[i].Value != expected[i] {
			t.Fatalf("expected %v, got %v", expected[i], got[i].Value)
		}
	}
}

func TestFind(t *testing.T) {
	l := NewList[int]()
	l.Push(4)
	l.Push(3)
	l.Push(1)
	l.Push(2)

	expected := []int{4, 3, 1, 2}
	for i := 0; i < 4; i++ {
		got := *l.Find(i)
		if got.Value != expected[i] {
			t.Fatalf("expected %v, got %v", expected[i], got.Value)
		}
	}
}
