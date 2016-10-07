// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package redblacktree implements a red-black tree.
//
// Used by TreeSet and TreeMap.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Red%E2%80%93black_tree
package redblacktree

import (
	"fmt"
	"github.com/emirpasic/gods/trees"
	"github.com/emirpasic/gods/utils"
)

func assertTreeImplementation() {
	var _ trees.Tree = (*Tree)(nil)
}

type color bool

const (
	black, red color = true, false
)

// Tree holds elements of the red-black tree
type Tree struct {
	Root       *Node
	size       int
	Comparator utils.Comparator
}

// Node is a single element within the tree
type Node struct {
	Key    interface{}
	Value  interface{}
	color  color
	Left   *Node
	Right  *Node
	Parent *Node
	Size int
}

// NewWith instantiates a red-black tree with the custom comparator.
func NewWith(comparator utils.Comparator) *Tree {
	return &Tree{Comparator: comparator}
}

// NewWithIntComparator instantiates a red-black tree with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator() *Tree {
	return &Tree{Comparator: utils.IntComparator}
}

// NewWithStringComparator instantiates a red-black tree with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator() *Tree {
	return &Tree{Comparator: utils.StringComparator}
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Get(key interface{}) (value interface{}, found bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.Value, true
	}
	return nil, false
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion,
// otherwise method panics.
func (tree *Tree) Put(key interface{}, value interface{}) {
	tree.Root = put(tree, tree.Root, key, value)
	tree.size = size(tree.Root)
	// always color the Root black
	if isRed(tree.Root) {
		tree.Root.color = black
	}
}

func put(tree *Tree, node *Node, key interface{}, value interface{}) *Node{
	if node == nil {
		return &Node{Key: key, Value: value, color: red, Size: 1}
	}
	compare := tree.Comparator(key, node.Key)
	if compare == 0 {
		node.Value = value
		return node
	} else if compare < 0 {
		node.Left = put(tree, node.Left, key, value)
	} else {
		node.Right = put(tree, node.Right, key, value)
	}
	return fixUp(node)
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Remove(key interface{}) {
	if (tree.Root == nil) {
		return
	}
	tree.Root.color = red
	tree.Root = remove(tree, tree.Root, key)
	if (tree.Root != nil) {
		tree.Root.color = black
	}
	tree.size = size(tree.Root)
}

func remove(tree *Tree, node *Node, key interface{}) *Node {
	if (node == nil) {
		return nil
	}
	//fmt.Println("Entering delete for node", node.Key, tree.String())
	if (tree.Comparator(key, node.Key) < 0) {
		// go Left
		if (!isRed(node.Left) && !isRed(node.Left.Left)) {
			node = moveRedLeft(node)
		}
		node.Left = remove(tree, node.Left, key)
	} else {
		// go Right
		if (isRed(node.Left)) {
			node = rotateRight(node)
		}
		if (tree.Comparator(key, node.Key) == 0 && node.Right == nil) {
			return nil
		}
		if (node.Right != nil && !isRed(node.Right) && !isRed(node.Right.Left)) {
			node = moveRedRight(node)
		}
		if (tree.Comparator(key, node.Key) == 0) {
			node = replaceNodeWithMin(node)
		} else {
			node.Right = remove(tree, node.Right, key)
		}
	}
	return fixUp(node)
}

// delete the key-value pair with the minimum key rooted at node
func removeMin(node *Node) *Node {
	if (node.Left == nil) {
		return nil
	}
	if (!isRed(node.Left) && !isRed(node.Left.Left)) {
		node = moveRedLeft(node)
	}
	node.Left = removeMin(node.Left);
	return fixUp(node);
}

// replace the node with minimum of the Tree rooted at the subTree
func replaceNodeWithMin(node *Node) *Node {
	subTreeMin := getMin(node.Right)
	rightChild := removeMin(node.Right)
	subTreeMin.Left = node.Left
	subTreeMin.Right = rightChild
	return subTreeMin
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
	keys := make([]interface{}, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

// Values returns all values in-order based on the key.
func (tree *Tree) Values() []interface{} {
	values := make([]interface{}, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

// Left returns the left-most (min) node or nil if tree is empty.
func (tree *Tree) Left() *Node {
	if (tree.Root == nil) {
		return nil
	}
	return getMin(tree.Root)
}

// Right returns the right-most (max) node or nil if tree is empty.
func (tree *Tree) Right() *Node {
	if (tree.Root == nil) {
		return nil
	}
	return getMax(tree.Root)
}

// Floor Finds floor node of the input key, return the floor node or nil if no ceiling is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree is larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Floor(key interface{}) (floor *Node, found bool) {
	found = false
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			node = node.Left
		case compare > 0:
			floor, found = node, true
			node = node.Right
		}
	}
	if found {
		return floor, true
	}
	return nil, false
}

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree is smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Ceiling(key interface{}) (ceiling *Node, found bool) {
	found = false
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			ceiling, found = node, true
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	if found {
		return ceiling, true
	}
	return nil, false
}

// Clear removes all nodes from the tree.
func (tree *Tree) Clear() {
	clearNode(tree.Root)
	tree.size = 0
}

// Helper to ensure that all the nodes are disconnected to allow for GC
func clearNode(node *Node) {
	if (node != nil) {
		node.Parent = nil
		clearNode(node.Left)
		clearNode(node.Right)
	}
}

// String returns a string representation of container
func (tree *Tree) String() string {
	str := "RedBlackTree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str)
	}
	return str
}

func (node *Node) String() string {
	if isRed(node) {
		return fmt.Sprintf("(%v)", node.Key)
	}
	return fmt.Sprintf("%v", node.Key)
}

func output(node *Node, prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}

func (tree *Tree) lookup(key interface{}) *Node {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
}

// Rotate the subtree rooted at node to the left and return the new root.
// Will make the right child of node the new root and will also update the
// parent pointers accordingly. The new root will have the same color as the
// old one and the left child of the root will be colored RED.
func rotateLeft(node *Node) *Node {
	newRoot := node.Right
	// rotate
	node.Right = newRoot.Left
	newRoot.Left = node
	// fix Parents
	newRoot.Parent = node.Parent
	node.Parent = newRoot
	if node.Right != nil {
		node.Right.Parent = node
	}
	// fix colors
	newRoot.color = newRoot.Left.color
	newRoot.Left.color = red
	// fix sizes
	node.Size = size(node)
	newRoot.Size = size(newRoot)
	return newRoot
}

// Rotate the subtree rooted at node to the right and return the new root.
// Will make the left child of node the new root and will update the parent
// pointers. The new root will be colored like the old one and the right child
// will be colored RED.
func rotateRight(node *Node) *Node {
	newRoot := node.Left
	// rotate
	node.Left = newRoot.Right
	newRoot.Right = node
	// fix Parents
	newRoot.Parent = node.Parent
	node.Parent = newRoot
	if node.Left != nil {
		node.Left.Parent = node
	}
	// fix colors
	newRoot.color = newRoot.Right.color
	newRoot.Right.color = red
	// fix sizes
	node.Size = size(node)
	newRoot.Size = size(newRoot)
	return newRoot
}

func moveRedLeft(node *Node) *Node {
	flipColors(node)
	if (node.Right != nil && isRed(node.Right.Left)) {
		node.Right = rotateRight(node.Right)
		node = rotateLeft(node)
		flipColors(node)
	}
	return node
}

func moveRedRight(node *Node) *Node {
	flipColors(node)
	if (isRed(node.Left.Left)) {
		node = rotateRight(node)
		flipColors(node)
	}
	return node
}

func fixUp(node *Node) *Node {
	if (isRed(node.Right)) {
		node = rotateLeft(node)
	}
	if (isRed(node.Left) && isRed(node.Left.Left)) {
		node = rotateRight(node)
	}
	if (childrenAreRed(node)) {
		flipColors(node)
	}
	node.Size = size(node.Left) + size(node.Right) + 1;
	return node;
}

// All valid LLRBs must have two properties:
// 1. All paths from the root to the leaves must have the same number of black
// nodes and
// 2. there must never be two consecutive red nodes.
func (tree *Tree) Validate() {
	if (tree.Root != nil) {
		countBlackNodes(tree.Root)
	}
}

func countBlackNodes(node *Node) int {
	if (node == nil) {
		return 0
	}
	count1 := countBlackNodes(node.Left)
	count2 := countBlackNodes(node.Right)
	// There is a different amount of black links to leaves from this node
	if (count1 != count2) {
		panic("Subtree rooted at node " + node.String() + " does not have an " +
			"equal number of black nodes to all leaves")
	}
	// There are two consecutive red links
	if (isRed(node) && (isRed(node.Right) || isRed(node.Left))) {
		panic("There are two consecutive links starting from node " + node.String())
	}
	// red nodes are not counted only black ones.
	if (isRed(node)) {
		return count1
	}
	return count1 + 1
}

////////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////////

// Returns whether or not the node is red. If the node is null it will return
// false.
func isRed(node *Node) bool {
	if node == nil {
		return false
	}
	return node.color == red
}

// Will return true iff subRoot is black and both children are non-nil and red.
func childrenAreRed(subRoot *Node) bool {
	// If the root or any of the children are nil return false
	if (subRoot == nil || subRoot.Left == nil || subRoot.Right == nil) {
		return false
	}
	return !isRed(subRoot) && isRed(subRoot.Left) && isRed(subRoot.Right)
}

// Flip the colors the subroot and the children.
// Calling function must make sure that node has two non-nil children.
// Also the two children must have the opposite color of the root.
func flipColors(node *Node) {
	if isRed(node) {
		node.color = black
	} else {
		node.color = red
	}
	if (node.Left != nil) {
		node.Left.color = !node.color
	}
	if (node.Right != nil) {
		node.Right.color = !node.color
	}
}

// Returns the size of the subtree rooted at this node.
// It will re-calculate the size of the subtree using recursion so it should be
// used only when the size of the tree changes.
func size(node *Node) int {
	if node == nil {
		return 0
	}
	return size(node.Left) + size(node.Right) + 1
}

func getMin(node *Node) *Node {
	if (node.Left == nil) {
		return node
	}
	return getMin(node.Left)
}

func getMax(node *Node) *Node {
	if (node.Right == nil) {
		return node
	}
	return getMax(node.Right)
}
