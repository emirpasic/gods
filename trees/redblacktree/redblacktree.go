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

	var subRoot *Node = node
	if isRed(subRoot.Right) && !isRed(subRoot.Left){
		subRoot = rotateLeft(subRoot)
	}
	if isRed(subRoot.Left) && isRed(subRoot.Left.Left) {
		subRoot = rotateRight(subRoot)
	}
	if (childrenAreRed(subRoot)) {
		flipColors(subRoot)
	}
	return subRoot
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Remove(key interface{}) {
	var child *Node
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == black {
			node.color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
	tree.size--
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
	var parent *Node
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Left
	}
	return parent
}

// Right returns the right-most (max) node or nil if tree is empty.
func (tree *Tree) Right() *Node {
	var parent *Node
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
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

func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
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

func (tree *Tree) replaceNode(old *Node, new *Node) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (node *Node) maximumNode() *Node {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tree *Tree) deleteCase1(node *Node) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *Tree) deleteCase2(node *Node) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.Parent.color = red
		sibling.color = black
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *Tree) deleteCase3(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *Tree) deleteCase4(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		node.Parent.color = black
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *Tree) deleteCase5(node *Node) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == red &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Right) == red &&
		nodeColor(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *Tree) deleteCase6(node *Node) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.Parent)
	node.Parent.color = black
	if node == node.Parent.Left && nodeColor(sibling.Right) == red {
		sibling.Right.color = black
		tree.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == red {
		sibling.Left.color = black
		tree.rotateRight(node.Parent)
	}
}

func nodeColor(node *Node) color {
	if node == nil {
		return black
	}
	return node.color
}

// All valid LLRBs must have two properties. All paths from the root to the
// leaves must have the same number of black nodes and there must never be
// two consecutive red nodes.
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
