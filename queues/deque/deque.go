package deque

// Deque represents a double-ended queue.
type Deque[T comparable] struct {
	items []T
}

// NewDeque creates a new Deque.
func NewDeque[T comparable]() *Deque[T] {
	return &Deque[T]{items: []T{}}
}

// AddFront adds an item to the front of the deque.
func (d *Deque[T]) AddFront(item T) {
	d.items = append([]T{item}, d.items...)
}

// AddBack adds an item to the back of the deque.
func (d *Deque[T]) AddBack(item T) {
	d.items = append(d.items, item)
}

// RemoveFront removes and returns the item from the front of the deque.
func (d *Deque[T]) RemoveFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

// RemoveBack removes and returns the item from the back of the deque.
func (d *Deque[T]) RemoveBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, true
}

// IsEmpty checks if the deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return len(d.items) == 0
}

// Size returns the number of items in the deque.
func (d *Deque[T]) Size() int {
	return len(d.items)
}

// withinRange checks if the index is within the range of the deque.
func (d *Deque[T]) withinRange(index int) bool {
	return index >= 0 && index < len(d.items)
}
