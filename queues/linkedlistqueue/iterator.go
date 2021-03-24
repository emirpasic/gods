package linkedlistqueue

import "github.com/emirpasic/gods/containers"

func assertIteratorImplementation() {
	var _ containers.IteratorWithIndex = (*Iterator)(nil)
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
type Iterator struct {
	queue *Queue
	index int
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (queue *Queue) Iterator() Iterator {
	return Iterator{queue: queue, index: -1}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.queue.Size() {
		iterator.index++
	}
	return iterator.queue.withinRange(iterator.index)
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	value, _ := iterator.queue.list.Get(iterator.index) // FIFO
	return value
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Begin() {
	iterator.index = -1
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}
