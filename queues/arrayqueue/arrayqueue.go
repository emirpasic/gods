package arrayqueue

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
)

// Queue holds elements in an array-list
type Queue struct {
	list *arraylist.List
}

// New instantiates a new empty queue
func New() *Queue {
	return &Queue{list: arraylist.New()}
}

// Push adds a value onto the back of the queue
func (queue *Queue) Push(value interface{}) {
	queue.list.Add(value)
}

// Pop removes front element in queue and returns it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to pop.
func (queue *Queue) Pop() (value interface{}, ok bool) {
	value, ok = queue.list.Get(0)
	queue.list.RemoveFirstElem()
	return
}

// Peek returns front element in the queue without removing it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to peek.
func (queue *Queue) Peek() (value interface{}, ok bool) {
	return queue.list.Get(0)
}

// Empty returns true if queue does not contain any elements.
func (queue *Queue) Empty() bool {
	return queue.list.Empty()
}

// Size returns number of elements within the queue.
func (queue *Queue) Size() int {
	return queue.list.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue) Clear() {
	queue.list.Clear()
}

// Values returns all elements in the queue (FIFO order).
func (queue *Queue) Values() []interface{} {
	size := queue.list.Size()
	elements := make([]interface{}, size, size)
	for i := 0; i < size; i++ {
		elements[i], _ = queue.list.Get(i) // FIFO
	}
	return elements
}

// String returns a string representation of container
func (queue *Queue) String() string {
	str := "ArrayQueue\n"
	values := []string{}
	for _, value := range queue.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (queue *Queue) withinRange(index int) bool {
	return index >= 0 && index < queue.list.Size()
}
