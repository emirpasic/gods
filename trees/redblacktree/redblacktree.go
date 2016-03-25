/*
Copyright (c) 2015, Emir Pasic
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Implementation of Red-black tree.
// Used by TreeSet and TreeMap.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Red%E2%80%93black_tree

package redblacktree

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/emirpasic/gods/trees"
	"github.com/emirpasic/gods/utils"
)

func assertInterfaceImplementation() {
	var _ trees.Interface = (*Tree)(nil)
}

type color bool

const (
	black, red color = true, false
)

type Tree struct {
	Root       *Node
	size       int
	comparator utils.Comparator
}

type Node struct {
	Key    interface{}
	Value  interface{}
	color  color
	Left   *Node
	Right  *Node
	Parent *Node
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
	insertedNode := &Node{Key: key, Value: value, color: red}
	if tree.Root == nil {
		tree.Root = insertedNode
	} else {
		node := tree.Root
		loop := true
		for loop {
			compare := tree.comparator(key, node.Key)
			switch {
			case compare == 0:
				node.Value = value
				return
			case compare < 0:
				if node.Left == nil {
					node.Left = insertedNode
					loop = false
				} else {
					node = node.Left
				}
			case compare > 0:
				if node.Right == nil {
					node.Right = insertedNode
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		insertedNode.Parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size += 1
}

// Searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Get(key interface{}) (value interface{}, found bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.Value, true
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
	tree.size -= 1
}

// Returns true if tree does not contain any nodes
func (tree *Tree) Empty() bool {
	return tree.size == 0
}

// Returns number of nodes in the tree.
func (tree *Tree) Size() int {
	return tree.size
}

// Returns all keys in-order
func (tree *Tree) Keys() []interface{} {
	keys := make([]interface{}, tree.size)
	for i, node := range tree.inOrder() {
		keys[i] = node.Key
	}
	return keys
}

// Returns all values in-order based on the key.
func (tree *Tree) Values() []interface{} {
	values := make([]interface{}, tree.size)
	for i, node := range tree.inOrder() {
		values[i] = node.Value
	}
	return values
}

// Removes all nodes from the tree.
func (tree *Tree) Clear() {
	tree.Root = nil
	tree.size = 0
}

func (tree *Tree) String() string {
	str := "RedBlackTree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str)
	}
	return str
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.Key)
}

func (tree *Tree) getMinFromNode(node *node) (foundNode *node, found bool) {
	if node.left == nil {
		return node, true
	} else {
		return tree.getMinFromNode(node.left)
	}
}

func (tree *Tree) GetMin() (value interface{}, found bool) {
	n, f := tree.getMinFromNode(tree.root)
	if n != nil {
		return n.value, f
	} else {
		return nil, false
	}
}

func (tree *Tree) getMaxFromNode(node *node) (foundNode *node, found bool) {
	if node.right == nil {
		return node, true
	} else {
		return tree.getMaxFromNode(node.right)
	}
}

func (tree *Tree) GetMax() (value interface{}, found bool) {
	n, f := tree.getMaxFromNode(tree.root)
	if n != nil {
		return n.value, f
	} else {
		return nil, false
	}
}

func (tree *Tree) RemoveMin() (value interface{}, deleted bool) {
	node, found := tree.getMinFromNode(tree.root)
	if found == true {
		tree.Remove(node.key)
		return node.value, found
	} else {
		return nil, false
	}
}

func (tree *Tree) RemoveMax() (value interface{}, deleted bool) {
	node, found := tree.getMaxFromNode(tree.root)
	if found == true {
		tree.Remove(node.key)
		return node.value, found
	} else {
		return nil, false
	}
}

// Returns all nodes in order
func (tree *Tree) inOrder() []*Node {
	nodes := make([]*Node, tree.size)
	if tree.size > 0 {
		current := tree.Root
		stack := linkedliststack.New()
		done := false
		count := 0
		for !done {
			if current != nil {
				stack.Push(current)
				current = current.Left
			} else {
				if !stack.Empty() {
					currentPop, _ := stack.Pop()
					current = currentPop.(*Node)
					nodes[count] = current
					count += 1
					current = current.Right
				} else {
					done = true
				}
			}
		}
	}
	return nodes
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
		compare := tree.comparator(key, node.Key)
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

func (node *Node) grandparent() *Node {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *Node) uncle() *Node {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	} else {
		return node.Parent.Left
	}
}

func (tree *Tree) rotateLeft(node *Node) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *Tree) rotateRight(node *Node) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
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

func (tree *Tree) insertCase1(node *Node) {
	if node.Parent == nil {
		node.color = black
	} else {
		tree.insertCase2(node)
	}
}

func (tree *Tree) insertCase2(node *Node) {
	if nodeColor(node.Parent) == black {
		return
	}
	tree.insertCase3(node)
}

func (tree *Tree) insertCase3(node *Node) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandparent().color = red
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *Tree) insertCase4(node *Node) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.insertCase5(node)
}

func (tree *Tree) insertCase5(node *Node) {
	node.Parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
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
	} else {
		tree.deleteCase2(node)
	}
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
