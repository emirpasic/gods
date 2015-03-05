/*
Copyright (c) Emir Pasic, All rights reserved.

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3.0 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library. See the file LICENSE included
with this distribution for more information.
*/

// Implementation of Red-black tree.
// Used by TreeSet and TreeMap.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Red%E2%80%93black_tree
package redblacktree

import (
	"fmt"
	"github.com/emirpasic/gods/utils"
)

type Color bool

const (
	BLACK, RED Color = true, false
)

type Tree struct {
	root       *Node
	size       int
	comparator utils.Comparator
}

type Node struct {
	key    interface{}
	value  interface{}
	color  Color
	left   *Node
	right  *Node
	parent *Node
}

// Instantiates a red-black tree with the custom comparator.
func NewWith(comparator utils.Comparator) *Tree {
	return &Tree{comparator: comparator}
}

// Instantiates a red-black tree with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator() *Tree {
	return &Tree{comparator: utils.IntComparator}
}

// Instantiates a red-black tree with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator() *Tree {
	return &Tree{comparator: utils.StringComparator}
}

// Inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Put(key interface{}, value interface{}) {
	insertedNode := &Node{key: key, value: value, color: RED}
	if tree.root == nil {
		tree.root = insertedNode
	} else {
		node := tree.root
		loop := true
		for loop {
			compare := tree.comparator(key, node.key)
			switch {
			case compare == 0:
				node.value = value
				return
			case compare < 0:
				if node.left == nil {
					node.left = insertedNode
					loop = false
				} else {
					node = node.left
				}
			case compare > 0:
				if node.right == nil {
					node.right = insertedNode
					loop = false
				} else {
					node = node.right
				}
			}
		}
		insertedNode.parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size += 1
}

// Searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Get(key interface{}) (interface{}, bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.value, true
	}
	return nil, false
}

// Remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Remove(key interface{}) {
	var child *Node
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.left != nil && node.right != nil {
		pred := node.left.maximumNode()
		node.key = pred.key
		node.value = pred.value
		node = pred
	}
	if node.left == nil || node.right == nil {
		if node.right == nil {
			child = node.left
		} else {
			child = node.right
		}
		if node.color == BLACK {
			node.color = color(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.parent == nil && child != nil {
			child.color = BLACK
		}
	}
	tree.size -= 1
}

// Returns true if tree does not contain any nodes
func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}

// Returns number of nodes in the tree.
func (tree *Tree) Size() int {
	return tree.size
}

func (tree *Tree) String() string {
	str := "RedBlackTree\n"
	if !tree.IsEmpty() {
		output(tree.root, "", true, &str)
	}
	return str
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.key)
}

func output(node *Node, prefix string, isTail bool, str *string) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.left, newPrefix, true, str)
	}
}

func (tree *Tree) lookup(key interface{}) *Node {
	node := tree.root
	for node != nil {
		compare := tree.comparator(key, node.key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.left
		case compare > 0:
			node = node.right
		}
	}
	return nil
}

func (node *Node) grandparent() *Node {
	if node != nil && node.parent != nil {
		return node.parent.parent
	}
	return nil
}

func (node *Node) uncle() *Node {
	if node == nil || node.parent == nil || node.parent.parent == nil {
		return nil
	}
	return node.parent.sibling()
}

func (node *Node) sibling() *Node {
	if node == nil || node.parent == nil {
		return nil
	}
	if node == node.parent.left {
		return node.parent.right
	} else {
		return node.parent.left
	}
}

func (tree *Tree) rotateLeft(node *Node) {
	right := node.right
	tree.replaceNode(node, right)
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.left = node
	node.parent = right
}

func (tree *Tree) rotateRight(node *Node) {
	left := node.left
	tree.replaceNode(node, left)
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.right = node
	node.parent = left
}

func (tree *Tree) replaceNode(old *Node, new *Node) {
	if old.parent == nil {
		tree.root = new
	} else {
		if old == old.parent.left {
			old.parent.left = new
		} else {
			old.parent.right = new
		}
	}
	if new != nil {
		new.parent = old.parent
	}
}

func (tree *Tree) insertCase1(node *Node) {
	if node.parent == nil {
		node.color = BLACK
	} else {
		tree.insertCase2(node)
	}
}

func (tree *Tree) insertCase2(node *Node) {
	if color(node.parent) == BLACK {
		return
	}
	tree.insertCase3(node)
}

func (tree *Tree) insertCase3(node *Node) {
	uncle := node.uncle()
	if color(uncle) == RED {
		node.parent.color = BLACK
		uncle.color = BLACK
		node.grandparent().color = RED
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *Tree) insertCase4(node *Node) {
	grandparent := node.grandparent()
	if node == node.parent.right && node.parent == grandparent.left {
		tree.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == grandparent.right {
		tree.rotateRight(node.parent)
		node = node.right
	}
	tree.insertCase5(node)
}

func (tree *Tree) insertCase5(node *Node) {
	node.parent.color = BLACK
	grandparent := node.grandparent()
	grandparent.color = RED
	if node == node.parent.left && node.parent == grandparent.left {
		tree.rotateRight(grandparent)
	} else if node == node.parent.right && node.parent == grandparent.right {
		tree.rotateLeft(grandparent)
	}
}

func (node *Node) maximumNode() *Node {
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	return node
}

func (tree *Tree) deleteCase1(node *Node) {
	if node.parent == nil {
		return
	} else {
		tree.deleteCase2(node)
	}
}

func (tree *Tree) deleteCase2(node *Node) {
	sibling := node.sibling()
	if color(sibling) == RED {
		node.parent.color = RED
		sibling.color = BLACK
		if node == node.parent.left {
			tree.rotateLeft(node.parent)
		} else {
			tree.rotateRight(node.parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *Tree) deleteCase3(node *Node) {
	sibling := node.sibling()
	if color(node.parent) == BLACK &&
		color(sibling) == BLACK &&
		color(sibling.left) == BLACK &&
		color(sibling.right) == BLACK {
		sibling.color = RED
		tree.deleteCase1(node.parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *Tree) deleteCase4(node *Node) {
	sibling := node.sibling()
	if color(node.parent) == RED &&
		color(sibling) == BLACK &&
		color(sibling.left) == BLACK &&
		color(sibling.right) == BLACK {
		sibling.color = RED
		node.parent.color = BLACK
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *Tree) deleteCase5(node *Node) {
	sibling := node.sibling()
	if node == node.parent.left &&
		color(sibling) == BLACK &&
		color(sibling.left) == RED &&
		color(sibling.right) == BLACK {
		sibling.color = RED
		sibling.left.color = BLACK
		tree.rotateRight(sibling)
	} else if node == node.parent.right &&
		color(sibling) == BLACK &&
		color(sibling.right) == RED &&
		color(sibling.left) == BLACK {
		sibling.color = RED
		sibling.right.color = BLACK
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *Tree) deleteCase6(node *Node) {
	sibling := node.sibling()
	sibling.color = color(node.parent)
	node.parent.color = BLACK
	if node == node.parent.left && color(sibling.right) == RED {
		sibling.right.color = BLACK
		tree.rotateLeft(node.parent)
	} else if color(sibling.left) == RED {
		sibling.left.color = BLACK
		tree.rotateRight(node.parent)
	}
}

func color(node *Node) Color {
	if node == nil {
		return BLACK
	}
	return node.color
}
