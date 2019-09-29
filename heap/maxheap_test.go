package heap

import (
	"math/rand"
	"testing"
)

func TestIntMaxHeap(t *testing.T) {
	h := NewIntMaxHeap()
	for i := 0; i < 1000; i++ {
		h.Add(rand.Int())
	}

	prev := 0
	for !h.Empty() {
		cur, err := h.Poll()
		if err != nil {
			t.Fatal(err.Error())
		}

		if prev != 0 && cur.(int) > prev {
			t.Fatal("invariant violated", cur, prev)
		}
		prev = cur.(int)
	}
}

func BenchmarkIntMaxHeapAdd100k(b *testing.B) {
	b.StopTimer()
	h := NewIntMaxHeap()
	b.StartTimer()
	for i := 0; i < 100000; i++ {
		h.Add(rand.Int())
	}
}

func BenchmarkIntMaxHeapPoll100k(b *testing.B) {
	b.StopTimer()
	h := NewIntMaxHeap()
	for i := 0; i < 100000; i++ {
		h.Add(rand.Int())
	}

	b.StartTimer()
	for !h.Empty() {
		h.Poll()
	}
}
