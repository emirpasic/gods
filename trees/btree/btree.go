// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package btree implements a B tree.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/B-tree
package btree

import (
	"fmt"
	"github.com/emirpasic/gods/trees"
	"github.com/emirpasic/gods/utils"
)

func assertTreeImplementation() {
	var _ trees.Tree = (*Tree)(nil)
}

// Tree holds elements of the B-tree
type Tree struct {
	Root       *Node            // Root node
	Comparator utils.Comparator // Key comparator
	size       int              // Total number of keys in the tree
	m          int              // Knuth order (maximum number of children)
}

// Node is a single element within the tree
type Node struct {
	Parent   *Node
	Entries  []*Entry // Contained keys in node
	Children []*Node  // Children nodes
}

// Entry represents the key-value pair contained within nodes
type Entry struct {
	Key   interface{}
	Value interface{}
}

// NewWith instantiates a B-tree with the Knuth order (maximum number of children) and a custom key comparator.
func NewWith(order int, comparator utils.Comparator) *Tree {
	if order < 2 {
		panic("Invalid order, should be at least 2")
	}
	return &Tree{m: order, Comparator: comparator}
}

// NewWithIntComparator instantiates a B-tree with the Knuth order (maximum number of children) and the IntComparator, i.e. keys are of type int.
func NewWithIntComparator(order int) *Tree {
	return NewWith(order, utils.IntComparator)
}

// NewWithStringComparator instantiates a B-tree with the Knuth order (maximum number of children) and the StringComparator, i.e. keys are of type string.
func NewWithStringComparator(order int) *Tree {
	return NewWith(order, utils.StringComparator)
}

// Put inserts key-value pair node into the tree.
// If key already exists, then its value is updated with the new value.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Put(key interface{}, value interface{}) {
	entry := &Entry{Key: key, Value: value}

	if tree.Root == nil {
		tree.Root = &Node{Entries: []*Entry{entry}, Children: []*Node{}}
		tree.size++
		return
	}

	if tree.insert(tree.Root, entry) {
		tree.size++
	}
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Get(key interface{}) (value interface{}, found bool) {
	return nil, false
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Remove(key interface{}) {
	// TODO
}

// Empty returns true if tree does not contain any nodes
func (tree *Tree) Empty() bool {
	return tree.size == 0
}

// Size returns number of nodes in the tree.
func (tree *Tree) Size() int {
	return tree.size
}

// Keys returns all keys in-order
func (tree *Tree) Keys() []interface{} {
	return nil // TODO
}

// Values returns all values in-order based on the key.
func (tree *Tree) Values() []interface{} {
	return nil // TODO
}

// Clear removes all nodes from the tree.
func (tree *Tree) Clear() {
	tree.Root = nil
	tree.size = 0
}

// String returns a string representation of container
func (tree *Tree) String() string {
	str := "BTree\n"
	if !tree.Empty() {
		str += tree.Root.String()
	}
	return str
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.Entries)
}

func (entry *Entry) String() string {
	return fmt.Sprintf("%v", entry.Key)
}

func (tree *Tree) isLeaf(node *Node) bool {
	return len(node.Children) == 0
}

func (tree *Tree) isFull(node *Node) bool {
	return len(node.Entries) == tree.maxEntries()
}

func (tree *Tree) shouldSplit(node *Node) bool {
	return len(node.Entries) > tree.maxEntries()
}

func (tree *Tree) maxChildren() int {
	return tree.m
}

func (tree *Tree) maxEntries() int {
	return tree.m - 1
}

func (tree *Tree) middle() int {
	return (tree.m - 1) / 2 // "-1" to favor right nodes to have more keys when splitting
}

func (tree *Tree) search(node *Node, entry *Entry) (index int, found bool) {
	low, high := 0, len(node.Entries)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		compare := tree.Comparator(entry.Key, node.Entries[mid].Key)
		switch {
		case compare > 0:
			low = mid + 1
		case compare < 0:
			high = mid - 1
		case compare == 0:
			return mid, true
		}
	}
	return low, false
}

func (tree *Tree) insert(node *Node, entry *Entry) (inserted bool) {
	if tree.isLeaf(node) {
		return tree.insertIntoLeaf(node, entry)
	}
	return tree.insertIntoInternal(node, entry)
}

func (tree *Tree) insertIntoLeaf(node *Node, entry *Entry) (inserted bool) {
	insertPosition, found := tree.search(node, entry)
	if found {
		node.Entries[insertPosition] = nil // GC
		node.Entries[insertPosition] = entry
		return false
	}
	node.Entries = append(node.Entries, nil)
	copy(node.Entries[insertPosition+1:], node.Entries[insertPosition:])
	node.Entries[insertPosition] = entry
	tree.split(node)
	return true
}

func (tree *Tree) insertIntoInternal(node *Node, entry *Entry) (inserted bool) {
	insertPosition, found := tree.search(node, entry)
	if found {
		node.Entries[insertPosition] = nil // GC
		node.Entries[insertPosition] = entry
		return false
	}
	return tree.insert(node.Children[insertPosition], entry)
}

func (tree *Tree) split(node *Node) {
	if !tree.shouldSplit(node) {
		return
	}

	if node == tree.Root {
		tree.splitRoot()
		return
	}

	tree.splitNonRoot(node)
}

func (tree *Tree) splitNonRoot(node *Node) {
	middle := tree.middle()
	parent := node.Parent

	left := &Node{Entries: node.Entries[:middle], Parent: parent}
	right := &Node{Entries: node.Entries[middle+1:], Parent: parent}

	if !tree.isLeaf(node) {
		left.Children = node.Children[:middle+1]
		right.Children = node.Children[middle+1:]
	}

	insertPosition, _ := tree.search(parent, node.Entries[middle])
	parent.Entries = append(parent.Entries, nil)
	copy(parent.Entries[insertPosition+1:], parent.Entries[insertPosition:])
	parent.Entries[insertPosition] = node.Entries[middle]

	parent.Children[insertPosition] = left

	parent.Children = append(parent.Children, nil)
	copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
	parent.Children[insertPosition+1] = right

	node = nil // GC

	tree.split(parent)
}

func (tree *Tree) splitRoot() {
	middle := tree.middle()

	left := &Node{Entries: tree.Root.Entries[:middle]}
	right := &Node{Entries: tree.Root.Entries[middle+1:]}

	if !tree.isLeaf(tree.Root) {
		left.Children = tree.Root.Children[:middle+1]
		right.Children = tree.Root.Children[middle+1:]
	}

	newRoot := &Node{
		Entries:  []*Entry{tree.Root.Entries[middle]},
		Children: []*Node{left, right},
	}

	left.Parent = newRoot
	right.Parent = newRoot
	tree.Root = newRoot
}
