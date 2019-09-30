package heap

import "errors"

// Heap is a specialized tree-based data structure which is essentially an almost complete tree that satisfies the heap
// property: in a max heap, for any given node C, if P is a parent node of C, then the key (the value) of P is
// greater than or equal to the key of C
type MinHeap struct {
	size    int
	heap    []interface{}
	greater func(h []interface{}, i int, j int) bool
}

// Creates Max Heap with a given greater function
func NewMinHeap(greater func(h []interface{}, i int, j int) bool) *MinHeap {
	return &MinHeap{
		heap:    make([]interface{}, 2),
		greater: greater,
	}
}

// Max Heap for strings
func NewStringMinHeap() *MinHeap {
	return NewMinHeap(func(h []interface{}, i int, j int) bool {
		return h[i].(string) > h[j].(string)
	})
}

// Max Heap for ints
func NewIntMinHeap() *MinHeap {
	return NewMinHeap(func(h []interface{}, i int, j int) bool {
		return h[i].(int) > h[j].(int)
	})
}

// Adds new element to heap with logarithmic amortized time
func (h *MinHeap) Add(value interface{}) {
	if h.size >= cap(h.heap)-1 {
		h.resize(h.size*2 + 1)
	}

	h.size++

	h.heap[h.size] = value
	h.swim(h.size)
}

// Polls element from heap with logarithmic amortized time
func (h *MinHeap) Poll() (interface{}, error) {
	if h.size == 0 {
		return nil, errors.New("heap is empty")
	}

	h.heap[1], h.heap[h.size] = h.heap[h.size], h.heap[1]
	min := h.heap[h.size]
	h.size--

	h.heap[h.size+1] = nil
	h.sink(1)

	if h.size < cap(h.heap)/2 {
		h.resize(h.size + 2)
	}

	return min, nil
}

// Returns max element without modifying heap
func (h *MinHeap) Peek() (interface{}, error) {
	if h.size == 0 {
		return nil, errors.New("heap is empty")
	}

	return h.heap[1], nil
}

// Returns size of this heap
func (h *MinHeap) Size() int {
	return h.size
}

//  Returns true if this heap
//  is empty.
func (h *MinHeap) Empty() bool {
	return h.size == 0
}

// Resizes heap`s internal array
func (h *MinHeap) resize(cap int) {
	t := make([]interface{}, cap)
	for i := 1; i <= h.size; i++ {
		t[i] = h.heap[i]
	}

	h.heap = t
}

func (h *MinHeap) swim(k int) {
	for k > 1 {
		if h.greater(h.heap, k/2, k) {
			h.heap[k], h.heap[k/2] = h.heap[k/2], h.heap[k]
		}
		k = k / 2
	}
}

func (h *MinHeap) sink(k int) {
	for k*2 < h.size {
		j := k * 2
		if h.greater(h.heap, j, j+1) {
			j = j + 1
		}

		if h.greater(h.heap, k, j) {
			h.heap[k], h.heap[j] = h.heap[j], h.heap[k]
		}

		k = j
	}
}
