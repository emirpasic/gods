package heap

import "errors"

// Heap is a specialized tree-based data structure which is essentially an almost complete tree that satisfies the heap
// property: in a max heap, for any given node C, if P is a parent node of C, then the key (the value) of P is
// greater than or equal to the key of C
type MaxHeap struct {
	size int
	heap []interface{}
	less func(h []interface{}, i int, j int) bool
}

// Creates Max Heap with a given less function
func NewMaxHeap(less func(h []interface{}, i int, j int) bool) *MaxHeap {
	return &MaxHeap{
		heap: make([]interface{}, 2),
		less: less,
	}
}

// Max Heap for strings
func NewStringMaxHeap() *MaxHeap {
	return NewMaxHeap(func(h []interface{}, i int, j int) bool {
		return h[i].(string) > h[j].(string)
	})
}

// Max Heap for ints
func NewIntMaxHeap() *MaxHeap {
	return NewMaxHeap(func(h []interface{}, i int, j int) bool {
		a := h[i].(int)
		b := h[j].(int)

		return a < b
	})
}

// Adds new element to heap with logarithmic amortized time
func (h *MaxHeap) Add(value interface{}) {
	if h.size >= cap(h.heap)-1 {
		h.resize(h.size*2 + 1)
	}

	h.size++

	h.heap[h.size] = value
	h.swim(h.size)
}

// Polls element from heap with logarithmic amortized time
func (h *MaxHeap) Poll() (interface{}, error) {
	if h.size == 0 {
		return nil, errors.New("heap is empty")
	}
	h.size--

	res := h.heap[1]
	h.heap[1] = h.heap[h.size+1]
	h.heap[h.size+1] = nil
	h.sink(1)

	if h.size < cap(h.heap)/2 {
		h.resize(h.size + 2)
	}

	return res, nil
}

// Returns max element without modifying heap
func (h *MaxHeap) Peek() (interface{}, error) {
	if h.size == 0 {
		return nil, errors.New("heap is empty")
	}

	return h.heap[1], nil
}

// Returns size of this heap
func (h *MaxHeap) Size() int {
	return h.size
}

//  Returns true if this heap
//  is empty.
func (h *MaxHeap) Empty() bool {
	return h.size == 0
}

// Resizes heap`s internal array
func (h *MaxHeap) resize(cap int) {
	t := make([]interface{}, cap)
	for i := 1; i <= h.size; i++ {
		t[i] = h.heap[i]
	}

	h.heap = t
}

func (h *MaxHeap) swim(k int) {
	for k > 1 {
		if h.less(h.heap, k/2, k) {
			h.heap[k], h.heap[k/2] = h.heap[k/2], h.heap[k]
		}
		k = k / 2
	}
}

func (h *MaxHeap) sink(k int) {
	for k*2 < h.size {
		j := k * 2
		if h.less(h.heap, j, j+1) {
			j = j + 1
		}

		if h.less(h.heap, k, j) {
			h.heap[k], h.heap[j] = h.heap[j], h.heap[k]
		}

		k = j
	}
}
