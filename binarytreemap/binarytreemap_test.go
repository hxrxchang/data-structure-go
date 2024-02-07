package binarytreemap

import "testing"

func Test(t *testing.T) {
	btm := NewBinaryTreeMap[int]()
	btm.Add("a", 1)
	btm.Add("b", 2)

	got, err := btm.Get("a")
	if err != nil || got != 1 {
		t.Errorf("want: %d, got: %v", 1, err)
	}

	got, err = btm.Get("b")
	if err != nil || got != 2 {
		t.Errorf("want: %d, got: %v", 2, err)
	}

	btm.Add("a", 3)
	got, err = btm.Get("a")
	if err != nil || got != 3 {
		t.Errorf("want: %d, got: %v", 3, err)
	}
}
