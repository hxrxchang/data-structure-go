package list

import "fmt"

type Element[T any] struct {
	Value T
	Next  *Element[T]
}

type List[T any] struct {
	First *Element[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Unshift(value T) {
	l.First = &Element[T]{
		Value: value,
		Next:  l.First,
	}
}

func (l *List[T]) Push(value T) {
	if l.First == nil {
		l.First = &Element[T]{
			Value: value,
			Next:  nil,
		}
	} else {
		var last *Element[T]
		for e := l.First; e != nil; e = e.Next {
			last = e
		}
		last.Next = &Element[T]{
			Value: value,
			Next:  nil,
		}
	}
}

func (l *List[T]) Each() []T {
	var result []T
	for e := l.First; e != nil; e = e.Next {
		result = append(result, e.Value)
	}
	return result
}

func (l *List[T]) Find(idx int) (T, error) {
	i := 0
	for e := l.First; e != nil; e = e.Next {
		if i == idx {
			return e.Value, nil
		}
		i++
	}
	var zero T
	return zero, fmt.Errorf("index out of range")
}
