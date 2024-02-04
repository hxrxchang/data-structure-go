package list

type Element[T any] struct {
	Value T
	Next *Element[T]
}

func NewElement[T any](value T, next *Element[T]) *Element[T] {
	return &Element[T]{Value: value, Next: next}
}


