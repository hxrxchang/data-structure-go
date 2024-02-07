package hashtable

import (
	"fmt"
	"hash/fnv"

	"github.com/hxrxchang/data-structure-go/list"
)

type HashTable[T any] struct {
	Bins []*list.List[KV[T]]
	Keys *list.List[string]
}

type KV[T any] struct {
	Key   string
	Value T
}

func NewHashTable[T any]() *HashTable[T] {
	binSize := 4
	var bins []*list.List[KV[T]]
	for i := 0; i < binSize; i++ {
		bins = append(bins, list.NewList[KV[T]]())
	}
	return &HashTable[T]{Bins: bins, Keys: list.NewList[string]()}
}

func (h *HashTable[T]) Set(key string, value T) {
	idx := hash(key) % len(h.Bins)
	bin := h.Bins[idx]
	// すでに同じキーがあれば上書き
	for e := bin.First; e != nil; e = e.Next {
		if e.Value.Key == key {
			e.Value.Value = value
			return
		}
	}
	bin.Unshift(KV[T]{Key: key, Value: value})
	// このListの実装では Unshift の方が計算量が少ないため Unshift を使う
	h.Keys.Unshift(key)
}

func (h *HashTable[T]) Get(key string) (T, error) {
	idx := hash(key) % len(h.Bins)
	bin := h.Bins[idx]
	for e := bin.First; e != nil; e = e.Next {
		if e.Value.Key == key {
			return e.Value.Value, nil
		}
	}
	var zero T
	return zero, fmt.Errorf("key: %s not found", key)
}

func (h *HashTable[T]) Each() []*KV[T] {
	var kvs []*KV[T]
	// unshiftで後から追加されたものが先頭に来ているので逆順にする
	keys := reverseSlice[string](h.Keys.Each())
	for _, key := range keys {
		value, _ := h.Get(key)
		kvs = append(kvs, &KV[T]{Key: key, Value: value})
	}
	return kvs
}

func hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())
}


func reverseSlice[T any](s []T) []T {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}
