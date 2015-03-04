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
// References: http://en.wikipedia.org/wiki/Red%E2%80%93black_tree
package redblacktree

import (
	"github.com/emirpasic/gods/utils"
)

type Color bool

const (
	BLACK, RED Color = true, false
)

type Tree struct {
	root       *Node
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

// Instantiates a red-black tree with the custom comparator
func NewWith(comparator utils.Comparator) *Tree {
	return &Tree{comparator: comparator}
}

// Instantiates a red-black tree with the IntComparator, i.e. keys are of type int
func NewWithIntComparator() *Tree {
	return &Tree{comparator: utils.IntComparator}
}

// Instantiates a red-black tree with the StringComparator, i.e. keys are of type string
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
					break
				} else {
					node = node.right
				}
			}
		}
		insertedNode.parent = node
	}
	tree.insertCase1(insertedNode)
}

// Searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics
func (tree *Tree) Get(key interface{}) (interface{}, bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.value, true
	}
	return nil, false
}

func (tree *Tree) Remove(key interface{}) {

}

// Returns true if tree does not contain any nodes
func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
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
	grandparent := node.grandparent()
	switch {
	case grandparent == nil:
		return nil
	case node.parent == grandparent.left:
		return grandparent.right
	default:
		return grandparent.left
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
	if node.parent.color == BLACK {
		return /* Tree is still valid */
	}
	tree.insertCase3(node)
}

func (tree *Tree) insertCase3(node *Node) {
	if node.uncle().color == RED {
		node.parent.color = BLACK
		node.uncle().color = BLACK
		node.grandparent().color = RED
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *Tree) insertCase4(node *Node) {
	if node == node.parent.right && node.parent == node.grandparent().left {
		tree.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == node.grandparent().right {
		tree.rotateRight(node.parent)
		node = node.right
	}
	tree.insertCase5(node)
}

func (tree *Tree) insertCase5(node *Node) {
	node.parent.color = BLACK
	node.grandparent().color = RED
	if node == node.parent.left && node.parent == node.grandparent().left {
		tree.rotateRight(node.grandparent())
	} else {
		tree.rotateLeft(node.grandparent())
	}
}
