package heap

import (
	"math/rand"
	"testing"
)

func TestIntMinHeap(t *testing.T) {
	h := NewIntMinHeap()
	//for i := 0; i < 1000; i++ {
	//	h.Add(rand.Int())
	//}

	h.Add(4)
	h.Add(6)
	h.Add(1)
	h.Add(2)
	h.Add(9)

	prev := 0
	for !h.Empty() {
		cur, err := h.Poll()
		if err != nil {
			t.Fatal(err.Error())
		}

		if prev != 0 && cur.(int) < prev {
			t.Fatal("invariant violated", cur, prev)
		}
		prev = cur.(int)
	}
}

func BenchmarkIntMinHeapAdd100k(b *testing.B) {
	b.StopTimer()
	h := NewIntMinHeap()
	b.StartTimer()
	for i := 0; i < 100000; i++ {
		h.Add(rand.Int())
	}
}

func BenchmarkIntMinHeapPoll100k(b *testing.B) {
	b.StopTimer()
	h := NewIntMinHeap()
	for i := 0; i < 100000; i++ {
		h.Add(rand.Int())
	}

	b.StartTimer()
	for !h.Empty() {
		h.Poll()
	}
}
