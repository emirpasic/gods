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
	root       *Node            // Root node
	comparator utils.Comparator // Key comparator
	size       int              // Total number of keys in the tree
	m          int              // Knuth order (maximum number of children)
}

// Node is a single element within the tree
type Node struct {
	parent   *Node
	entries  []*Entry // Contained keys in node
	children []*Node  // Children nodes
}

type Entry struct {
	key   interface{}
	value interface{}
}

// NewWith instantiates a B-tree with the Knuth order (maximum number of children) and a custom key comparator.
func NewWith(order int, comparator utils.Comparator) *Tree {
	if order < 2 {
		panic("Invalid order, should be at least 2")
	}
	return &Tree{m: order, comparator: comparator}
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
	entry := &Entry{key: key, value: value}

	if tree.root == nil {
		tree.root = &Node{entries: []*Entry{entry}, children: []*Node{}}
		tree.size++
		return
	}

	if tree.insert(tree.root, entry) {
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
	tree.root = nil
	tree.size = 0
}

// String returns a string representation of container
func (tree *Tree) String() string {
	str := "BTree\n"
	if !tree.Empty() {
		str += tree.root.String()
	}
	return str
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.entries)
}

func (entry *Entry) String() string {
	return fmt.Sprintf("%v", entry.key)
}

func (tree *Tree) isLeaf(node *Node) bool {
	return len(node.children) == 0
}

func (tree *Tree) isFull(node *Node) bool {
	return len(node.entries) == tree.maxEntries()
}

func (tree *Tree) shouldSplit(node *Node) bool {
	return len(node.entries) > tree.maxEntries()
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
	low, high := 0, len(node.entries) - 1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		compare := tree.comparator(entry.key, node.entries[mid].key)
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
		node.entries[insertPosition] = nil // GC
		node.entries[insertPosition] = entry
		return false
	}
	node.entries = append(node.entries, nil)
	copy(node.entries[insertPosition + 1:], node.entries[insertPosition:])
	node.entries[insertPosition] = entry
	tree.split(node)
	return true
}

func (tree *Tree) insertIntoInternal(node *Node, entry *Entry) (inserted bool) {
	insertPosition, found := tree.search(node, entry)
	if found {
		node.entries[insertPosition] = nil // GC
		node.entries[insertPosition] = entry
		return false
	}
	return tree.insert(node.children[insertPosition], entry)
}

func (tree *Tree) split(node *Node) {
	if !tree.shouldSplit(node) {
		return
	}

	if node == tree.root {
		tree.splitRoot()
		return
	}

	tree.splitNonRoot(node)
}

func (tree *Tree) splitNonRoot(node *Node) {
	middle := tree.middle()
	parent := node.parent

	left := &Node{entries: node.entries[:middle], parent: parent}
	right := &Node{entries: node.entries[middle + 1:], parent: parent}

	if !tree.isLeaf(node) {
		left.children = node.children[:middle + 1]
		right.children = node.children[middle + 1:]
	}

	insertPosition, _ := tree.search(parent, node.entries[middle])
	parent.entries = append(parent.entries, nil)
	copy(parent.entries[insertPosition + 1:], parent.entries[insertPosition:])
	parent.entries[insertPosition] = node.entries[middle]

	parent.children[insertPosition] = left

	parent.children = append(parent.children, nil)
	copy(parent.children[insertPosition + 2:], parent.children[insertPosition + 1:])
	parent.children[insertPosition + 1] = right

	node = nil // GC

	tree.split(parent)
}

func (tree *Tree) splitRoot() {
	middle := tree.middle()

	left := &Node{entries: tree.root.entries[:middle]}
	right := &Node{entries: tree.root.entries[middle + 1:]}

	if !tree.isLeaf(tree.root) {
		left.children = tree.root.children[:middle + 1]
		right.children = tree.root.children[middle + 1:]
	}

	newRoot := &Node{
		entries:  []*Entry{tree.root.entries[middle]},
		children: []*Node{left, right},
	}

	left.parent = newRoot
	right.parent = newRoot
	tree.root = newRoot
}
