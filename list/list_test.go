package list

import "testing"

func TestNewElement(t *testing.T) {
	ele1 := NewElement[int](1, nil)
	if ele1.Value != 1 {
		t.Errorf("ele.Value = %v; want 1", ele1.Value)
	}
	if ele1.Next != nil {
		t.Errorf("ele.Next = %v; want nil", ele1.Next)
	}

	ele2 := NewElement[int](2, ele1)
	if ele2.Value != 2 {
		t.Errorf("ele.Value = %v; want 2", ele2.Value)
	}
	if ele2.Next != ele1 {
		t.Errorf("ele.Next = %v; want %v", ele2.Next, ele1)
	}
}
