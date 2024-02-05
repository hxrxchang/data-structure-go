package list

type Element[T any] struct {
	Value T
	Next *Element[T]
}


type List[T any] struct {
	First *Element[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Unshift(value T) {
	l.First =  &Element[T]{
		Value: value,
		Next: l.First,
	}
}

func (l *List[T]) Push(value T) {
	if l.First == nil {
		l.First = &Element[T]{
			Value: value,
			Next: nil,
		}
	} else {
		var last *Element[T]
		for e := l.First; e != nil; e = e.Next {
			last = e
		}
		last.Next = &Element[T]{
			Value: value,
			Next: nil,
		}
	}
}

func (l *List[T]) Each() []*Element[T] {
	var result []*Element[T]
	for e := l.First; e != nil; e = e.Next {
		result = append(result, e)
	}
	return result
}

func (l *List[T]) Find(idx int) *Element[T] {
	i := 0
	for e := l.First; e != nil; e = e.Next {
		if i == idx {
			return e
		}
		i++
	}
	return nil
}
