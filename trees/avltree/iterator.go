// Copyright (c) 2017, Benjamin Scher Purcell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package avltree

import "github.com/emirpasic/gods/containers"

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithKey = (*Iterator)(nil)
}

// Iterator holding the iterator's state
type Iterator struct {
	tree     *Tree
	node     *Node
	position position
}

type position byte

const (
	begin, between, end position = 0, 1, 2
)

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (tree *Tree) Iterator() Iterator {
	return Iterator{tree: tree, node: nil, position: begin}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iter *Iterator) Next() bool {
	switch iter.position {
	case begin:
		iter.position = between
		iter.node = iter.tree.Left()
	case between:
		iter.node = iter.node.Next()
	}

	if iter.node == nil {
		iter.position = end
		return false
	}
	return true
}

// Prev moves the iterator to the next element and returns true if there was a previous element in the container.
// If Prev() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Prev() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iter *Iterator) Prev() bool {
	switch iter.position {
	case end:
		iter.position = between
		iter.node = iter.tree.Right()
	case between:
		iter.node = iter.node.Prev()
	}

	if iter.node == nil {
		iter.position = begin
		return false
	}
	return true
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iter *Iterator) Value() interface{} {
	return iter.node.Value
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iter *Iterator) Key() interface{} {
	return iter.node.Key
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iter *Iterator) Begin() {
	iter.node = nil
	iter.position = begin
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iter *Iterator) End() {
	iter.node = nil
	iter.position = end
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator
func (iter *Iterator) First() bool {
	iter.Begin()
	return iter.Next()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iter *Iterator) Last() bool {
	iter.End()
	return iter.Prev()
}
