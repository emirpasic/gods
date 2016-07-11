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
	"bytes"
	"fmt"
	"github.com/emirpasic/gods/trees"
	"github.com/emirpasic/gods/utils"
	"strings"
)

func assertTreeImplementation() {
	var _ trees.Tree = (*Tree)(nil)
}

// Tree holds elements of the B-tree
type Tree struct {
	Root       *Node            // Root node
	Comparator utils.Comparator // Key comparator
	size       int              // Total number of keys in the tree
	m          int              // order (maximum number of children)
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

// NewWith instantiates a B-tree with the order (maximum number of children) and a custom key comparator.
func NewWith(order int, comparator utils.Comparator) *Tree {
	if order < 3 {
		panic("Invalid order, should be at least 3")
	}
	return &Tree{m: order, Comparator: comparator}
}

// NewWithIntComparator instantiates a B-tree with the order (maximum number of children) and the IntComparator, i.e. keys are of type int.
func NewWithIntComparator(order int) *Tree {
	return NewWith(order, utils.IntComparator)
}

// NewWithStringComparator instantiates a B-tree with the order (maximum number of children) and the StringComparator, i.e. keys are of type string.
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
	if tree.Empty() {
		return nil, false
	}
	node := tree.Root
	for {
		index, found := tree.search(node, key)
		if found {
			return node.Entries[index].Value, true
		}
		if tree.isLeaf(node) {
			return nil, false
		}
		node = node.Children[index]
	}
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

// Height returns the height of the tree.
func (tree *Tree) Height() int {
	return tree.Root.height()
}

// Left returns the left-most (min) node or nil if tree is empty.
func (tree *Tree) Left() *Node {
	if tree.Empty() {
		return nil
	}
	node := tree.Root
	for {
		if tree.isLeaf(node) {
			return node
		}
		node = node.Children[0]
	}
}

// LeftKey returns the left-most (min) key or nil if tree is empty.
func (tree *Tree) LeftKey() interface{} {
	if left := tree.Left(); left != nil {
		return left.Entries[0].Key
	}
	return nil
}

// LeftValue returns the left-most value or nil if tree is empty.
func (tree *Tree) LeftValue() interface{} {
	if left := tree.Left(); left != nil {
		return left.Entries[0].Value
	}
	return nil
}

// Right returns the right-most (max) node or nil if tree is empty.
func (tree *Tree) Right() *Node {
	if tree.Empty() {
		return nil
	}
	node := tree.Root
	for {
		if tree.isLeaf(node) {
			return node
		}
		node = node.Children[len(node.Children)-1]
	}
}

// RightKey returns the right-most (max) key or nil if tree is empty.
func (tree *Tree) RightKey() interface{} {
	if right := tree.Right(); right != nil {
		return right.Entries[len(right.Entries)-1].Key
	}
	return nil
}

// RightValue returns the right-most value or nil if tree is empty.
func (tree *Tree) RightValue() interface{} {
	if right := tree.Right(); right != nil {
		return right.Entries[len(right.Entries)-1].Value
	}
	return nil
}

// String returns a string representation of container (for debugging purposes)
func (tree *Tree) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("BTree\n")
	if !tree.Empty() {
		tree.output(&buffer, tree.Root, 0, true)
	}
	return buffer.String()
}

func (entry *Entry) String() string {
	return fmt.Sprintf("%v", entry.Key)
}

func (tree *Tree) output(buffer *bytes.Buffer, node *Node, level int, isTail bool) {
	for e := 0; e < len(node.Entries)+1; e++ {
		if e < len(node.Children) {
			tree.output(buffer, node.Children[e], level+1, true)
		}
		if e < len(node.Entries) {
			buffer.WriteString(strings.Repeat("    ", level))
			buffer.WriteString(fmt.Sprintf("%v", node.Entries[e].Key) + "\n")
		}
	}
}

func (node *Node) height() int {
	height := 0
	for ; node != nil; node = node.Children[0] {
		height++
		if len(node.Children) == 0 {
			break
		}
	}
	return height
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

// search searches only within the single node among its entries
func (tree *Tree) search(node *Node, key interface{}) (index int, found bool) {
	low, high := 0, len(node.Entries)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		compare := tree.Comparator(key, node.Entries[mid].Key)
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
	insertPosition, found := tree.search(node, entry.Key)
	if found {
		node.Entries[insertPosition] = entry
		return false
	}
	// Insert entry's key in the middle of the node
	node.Entries = append(node.Entries, nil)
	copy(node.Entries[insertPosition+1:], node.Entries[insertPosition:])
	node.Entries[insertPosition] = entry
	tree.split(node)
	return true
}

func (tree *Tree) insertIntoInternal(node *Node, entry *Entry) (inserted bool) {
	insertPosition, found := tree.search(node, entry.Key)
	if found {
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

	left := &Node{Entries: append([]*Entry(nil), node.Entries[:middle]...), Parent: parent}
	right := &Node{Entries: append([]*Entry(nil), node.Entries[middle+1:]...), Parent: parent}

	// Move children from the node to be split into left and right nodes
	if !tree.isLeaf(node) {
		left.Children = append([]*Node(nil), node.Children[:middle+1]...)
		right.Children = append([]*Node(nil), node.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	insertPosition, _ := tree.search(parent, node.Entries[middle].Key)

	// Insert middle key into parent
	parent.Entries = append(parent.Entries, nil)
	copy(parent.Entries[insertPosition+1:], parent.Entries[insertPosition:])
	parent.Entries[insertPosition] = node.Entries[middle]

	// Set child left of inserted key in parent to the created left node
	parent.Children[insertPosition] = left

	// Set child right of inserted key in parent to the created right node
	parent.Children = append(parent.Children, nil)
	copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
	parent.Children[insertPosition+1] = right

	tree.split(parent)
}

func (tree *Tree) splitRoot() {
	middle := tree.middle()

	left := &Node{Entries: append([]*Entry(nil), tree.Root.Entries[:middle]...)}
	right := &Node{Entries: append([]*Entry(nil), tree.Root.Entries[middle+1:]...)}

	// Move children from the node to be split into left and right nodes
	if !tree.isLeaf(tree.Root) {
		left.Children = append([]*Node(nil), tree.Root.Children[:middle+1]...)
		right.Children = append([]*Node(nil), tree.Root.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	// Root is a node with one entry and two children (left and right)
	newRoot := &Node{
		Entries:  []*Entry{tree.Root.Entries[middle]},
		Children: []*Node{left, right},
	}

	left.Parent = newRoot
	right.Parent = newRoot
	tree.Root = newRoot
}

func setParent(nodes []*Node, parent *Node) {
	for _, node := range nodes {
		node.Parent = parent
	}
}
