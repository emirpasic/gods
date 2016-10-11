// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktree

import (
	"errors"

	"github.com/emirpasic/gods/containers"
)

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

func traverseParents(node *Node, toTheRight bool) *Node {
	switch {
	case node.Parent == nil:
		return nil
	case (node.Parent.Left == node && toTheRight) || (node.Parent.Right == node && !toTheRight):
		return node.Parent
	case (node.Parent.Right == node && toTheRight) || (node.Parent.Left == node && !toTheRight):
		return traverseParents(node.Parent, toTheRight)
	}
	return nil
}

func getLeftMost(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return getLeftMost(node.Left)
}

func getRightMost(node *Node) *Node {
	if node.Right == nil {
		return node
	}
	return getRightMost(node.Right)
}

// Next moves the iterator to the next element and returns true if there was a
// next element in the container.
// If Next() returns true, then next element's key and value can be retrieved
// by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to
// the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator) Next() bool {
	switch iterator.position {
	case begin:
		if iterator.tree.Empty() {
			iterator.position = end
			return false
		}
		iterator.node = iterator.tree.Left()
		iterator.position = between
		return true
	case between:
		var nextNode *Node
		if iterator.node.Right != nil {
			nextNode = getLeftMost(iterator.node.Right)
		} else {
			nextNode = traverseParents(iterator.node, true)
			if nextNode == nil {
				iterator.position = end
				return false
			}
		}
		iterator.node = nextNode
		return true
	}
	return false
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Prev() bool {
	switch iterator.position {
	case between:
		var nextNode *Node
		if iterator.node.Left != nil {
			nextNode = getRightMost(iterator.node.Left)
		} else {
			nextNode = traverseParents(iterator.node, false)
			if nextNode == nil {
				iterator.position = begin
				return false
			}
		}
		iterator.node = nextNode
		return true
	case end:
		if iterator.tree.Empty() {
			return false
		}
		iterator.position = between
		iterator.node = iterator.tree.Right()
		return true
	}
	return false
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	return iterator.node.Value
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iterator *Iterator) Key() interface{} {
	return iterator.node.Key
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Begin() {
	iterator.node = nil
	iterator.position = begin
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator) End() {
	iterator.node = nil
	iterator.position = end
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Last moves the iterator to the last element and returns true if there was a
// last element in the container. If Last() returns true, then last element's
// key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.
func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}

// RangedIterator is a special type of Iterator for LLRB that will iterate only
// over elements whose keys are within the range provided during initialization
type RangedIterator struct {
	iterator *Iterator
	lo       interface{}
	high     interface{}
}

// IteratorWithin will return an iterator that will iterate through the nodes
// with keys within the range provided. The range is inclusive.
func (tree *Tree) IteratorWithin(lo interface{}, high interface{}) (*RangedIterator, error) {
	if tree.Comparator(lo, high) >= 0 {
		return nil, errors.New("The lo value should be strictly less than the high value")
	}
	it := tree.Iterator()
	return &RangedIterator{
		iterator: &it,
		high:     high,
		lo:       lo}, nil
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *RangedIterator) Value() interface{} {
	return iterator.iterator.node.Value
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iterator *RangedIterator) Key() interface{} {
	return iterator.iterator.node.Key
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *RangedIterator) Begin() {
	iterator.iterator.node = nil
	iterator.iterator.position = begin
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *RangedIterator) End() {
	iterator.iterator.node = nil
	iterator.iterator.position = end
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator
func (iterator *RangedIterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Last moves the iterator to the last element and returns true if there was a
// last element in the container. If Last() returns true, then last element's
// key and value can be retrieved by Key() and Value(). Modifies the state of the iterator.
func (iterator *RangedIterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}

// Search in the tree for key equal to lo or a key as close as possible to it's
// value that is also less or equal to high
func exploreAndGetClosestLargerElement(tree *Tree, lo interface{}, high interface{}) *Node {
	node := tree.Root
	var possibleNode *Node
	var compare int
	for node != nil {
		compare = tree.Comparator(node.Key, lo)
		if compare == 0 {
			return node
		} else if compare < 0 {
			// we need to go right to find larger keys
			node = node.Right
		} else {
			// we need to go left to smaller keys
			if tree.Comparator(node.Key, high) <= 0 {
				// this key is within the range so we should mark it
				possibleNode = node
			}
			node = node.Left
		}
	}
	return possibleNode
}

// Search in the tree for key equal to lo or a key as close as possible to it's
// value that is also less or equal to high
func exploreAndGetClosestSmallerElement(tree *Tree, high interface{}, lo interface{}) *Node {
	node := tree.Root
	var possibleNode *Node
	for node != nil {
		compare := tree.Comparator(node.Key, high)
		if compare == 0 {
			return node
		} else if compare > 0 {
			// we need to go left to find smaller keys
			node = node.Left
		} else {
			// we need to go right to larger keys
			if tree.Comparator(lo, node.Key) <= 0 {
				// this key is within the range so we should mark it
				possibleNode = node
			}
			node = node.Right
		}
	}
	return possibleNode
}

// Next will move the iterator to the node of the tree whose key is
// immediately larger. If the key value is outside the range provided during
// the initialization of the RangedIterator instance false is returned and the
// iterator posistion is set to end. Also, if there are no more larger keys
// available again false will be returned and the iterator will be set to position
// end. If the iterator is already at position end then false will be returned
// again
func (iterator *RangedIterator) Next() bool {
	switch iterator.iterator.position {
	case begin:
		closestLo := exploreAndGetClosestLargerElement(iterator.iterator.tree,
			iterator.lo, iterator.high)
		if closestLo == nil {
			iterator.iterator.position = end
			return false
		}
		iterator.iterator.node = closestLo
		iterator.iterator.position = between
		return true
	case between:
		if !iterator.iterator.Next() {
			return false
		}
		if iterator.iterator.tree.Comparator(iterator.iterator.node.Key, iterator.high) > 0 {
			iterator.iterator.position = end
			iterator.iterator.node = nil
			return false
		}
		return true
	default:
		return false
	}
}

// Prev will move the iterator to the first node whose key is smaller than the
// current one. If no node is available, or if there is a node but the node's
// key is outside the range provided when initializing the iterator, then the
// iterator will be set to position being and false will be returned. If the
// iterator is already at the position begin, again false will be returned
func (iterator *RangedIterator) Prev() bool {
	switch iterator.iterator.position {
	case end:
		closestHigh := exploreAndGetClosestSmallerElement(iterator.iterator.tree,
			iterator.high, iterator.lo)
		if closestHigh == nil {
			iterator.iterator.position = begin
			return false
		}
		iterator.iterator.node = closestHigh
		iterator.iterator.position = between
		return true
	case between:
		if !iterator.iterator.Prev() {
			return false
		}
		if iterator.iterator.tree.Comparator(iterator.iterator.node.Key, iterator.lo) < 0 {
			iterator.iterator.position = begin
			iterator.iterator.node = nil
			return false
		}
		return true
	default:
		return false
	}
}
