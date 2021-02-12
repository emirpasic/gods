package trie

import "github.com/emirpasic/gods/containers"

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithIndex = (*Iterator)(nil)
}

// Iterator holding the iterator's state
type Iterator struct {
	words []string
	trie *Trie
	index int
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (trie *Trie) Iterator() Iterator {
	s := make([]string, 10)
	findWords(trie.root, "", s)
	return Iterator{words:s, trie: trie, index:-1}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.heap.Size() {
		iterator.index++
	}
	return iterator.words.withinRange(iterator.index)
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.words.withinRange(iterator.index)
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	value, _ := iterator.words.list.Get(iterator.index)
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

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator) End() {
	iterator.index = iterator.words.Size()
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}
