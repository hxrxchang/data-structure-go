package hashtable

import (
	"testing"
)

func Test(t *testing.T) {
	h := NewHashTable[int]()
	// 追加
	keys := []string{"a", "b", "c", "d"}
	values := []int{1, 2, 3, 4}
	for i := 0; i < 4; i++ {
		h.Set(keys[i], values[i])
	}

	// Get
	expected := []int{1, 2, 3, 4}
	for i := 0; i < 4; i++ {
		got, err := h.Get(keys[i])
		if err != nil || got != expected[i] {
			t.Fatalf("expected %v, got %v", expected[i], got)
		}
	}

	// Each
	kvs := h.Each()
	if len(kvs) != 4 {
		t.Fatalf("expected 4, got %v", len(kvs))
	}
	for i, v := range kvs {
		if v.Key != keys[i] || v.Value != values[i] {
			t.Fatalf("expected %v, got %v", &KV[int]{Key: keys[i], Value: values[i]}, v)
		}
	}


	// 上書き
	h.Set("a", 5)
	got, err := h.Get("a")
	if err != nil || got != 5 {
		t.Fatalf("expected 5, got %v", got)
	}

	// キーがない場合
	_, err = h.Get("e")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
