// Copyright (c) 2017, Benjamin Scher Purcell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package avltree implements an AVL balanced binary tree.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/AVL_tree
package avltree

import (
	"cmp"
	"fmt"

	"github.com/emirpasic/gods/v2/trees"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert Tree implementation
var _ trees.Tree[int] = (*Tree[string, int])(nil)

// Tree holds elements of the AVL tree.
type Tree[K comparable, V any] struct {
	Root       *Node[K, V]         // Root node
	Comparator utils.Comparator[K] // Key comparator
	size       int                 // Total number of keys in the tree
}

// Node is a single element within the tree
type Node[K comparable, V any] struct {
	Key      K
	Value    V
	Parent   *Node[K, V]    // Parent node
	Children [2]*Node[K, V] // Children nodes
	b        int8
}

// New instantiates an AVL tree with the built-in comparator for K
func New[K cmp.Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{Comparator: cmp.Compare[K]}
}

// NewWith instantiates an AVL tree with the custom comparator.
func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Tree[K, V] {
	return &Tree[K, V]{Comparator: comparator}
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[K, V]) Put(key K, value V) {
	tree.put(key, value, nil, &tree.Root)
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[K, V]) Get(key K) (value V, found bool) {
	n := tree.GetNode(key)
	if n != nil {
		return n.Value, true
	}
	return value, false
}

// GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[K, V]) GetNode(key K) *Node[K, V] {
	n := tree.Root
	for n != nil {
		cmp := tree.Comparator(key, n.Key)
		switch {
		case cmp == 0:
			return n
		case cmp < 0:
			n = n.Children[0]
		case cmp > 0:
			n = n.Children[1]
		}
	}
	return n
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[K, V]) Remove(key K) {
	tree.remove(key, &tree.Root)
}

// Empty returns true if tree does not contain any nodes.
func (tree *Tree[K, V]) Empty() bool {
	return tree.size == 0
}

// Size returns the number of elements stored in the tree.
func (tree *Tree[K, V]) Size() int {
	return tree.size
}

// Size returns the number of elements stored in the subtree.
// Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.
func (n *Node[K, V]) Size() int {
	if n == nil {
		return 0
	}
	size := 1
	if n.Children[0] != nil {
		size += n.Children[0].Size()
	}
	if n.Children[1] != nil {
		size += n.Children[1].Size()
	}
	return size
}

// Keys returns all keys in-order
func (tree *Tree[K, V]) Keys() []K {
	keys := make([]K, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

// Values returns all values in-order based on the key.
func (tree *Tree[K, V]) Values() []V {
	values := make([]V, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

// Left returns the minimum element of the AVL tree
// or nil if the tree is empty.
func (tree *Tree[K, V]) Left() *Node[K, V] {
	return tree.bottom(0)
}

// Right returns the maximum element of the AVL tree
// or nil if the tree is empty.
func (tree *Tree[K, V]) Right() *Node[K, V] {
	return tree.bottom(1)
}

// Floor Finds floor node of the input key, return the floor node or nil if no floor is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree is larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[K, V]) Floor(key K) (floor *Node[K, V], found bool) {
	found = false
	n := tree.Root
	for n != nil {
		c := tree.Comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			n = n.Children[0]
		case c > 0:
			floor, found = n, true
			n = n.Children[1]
		}
	}
	if found {
		return
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
func (tree *Tree[K, V]) Ceiling(key K) (floor *Node[K, V], found bool) {
	found = false
	n := tree.Root
	for n != nil {
		c := tree.Comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			floor, found = n, true
			n = n.Children[0]
		case c > 0:
			n = n.Children[1]
		}
	}
	if found {
		return
	}
	return nil, false
}

// Clear removes all nodes from the tree.
func (tree *Tree[K, V]) Clear() {
	tree.Root = nil
	tree.size = 0
}

// String returns a string representation of container
func (tree *Tree[K, V]) String() string {
	str := "AVLTree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str)
	}
	return str
}

func (n *Node[K, V]) String() string {
	return fmt.Sprintf("%v", n.Key)
}

func (tree *Tree[K, V]) put(key K, value V, p *Node[K, V], qp **Node[K, V]) bool {
	q := *qp
	if q == nil {
		tree.size++
		*qp = &Node[K, V]{Key: key, Value: value, Parent: p}
		return true
	}

	c := tree.Comparator(key, q.Key)
	if c == 0 {
		q.Key = key
		q.Value = value
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	var fix bool
	fix = tree.put(key, value, q, &q.Children[a])
	if fix {
		return putFix(int8(c), qp)
	}
	return false
}

func (tree *Tree[K, V]) remove(key K, qp **Node[K, V]) bool {
	q := *qp
	if q == nil {
		return false
	}

	c := tree.Comparator(key, q.Key)
	if c == 0 {
		tree.size--
		if q.Children[1] == nil {
			if q.Children[0] != nil {
				q.Children[0].Parent = q.Parent
			}
			*qp = q.Children[0]
			return true
		}
		fix := removeMin(&q.Children[1], &q.Key, &q.Value)
		if fix {
			return removeFix(-1, qp)
		}
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	fix := tree.remove(key, &q.Children[a])
	if fix {
		return removeFix(int8(-c), qp)
	}
	return false
}

func removeMin[K comparable, V any](qp **Node[K, V], minKey *K, minVal *V) bool {
	q := *qp
	if q.Children[0] == nil {
		*minKey = q.Key
		*minVal = q.Value
		if q.Children[1] != nil {
			q.Children[1].Parent = q.Parent
		}
		*qp = q.Children[1]
		return true
	}
	fix := removeMin(&q.Children[0], minKey, minVal)
	if fix {
		return removeFix(1, qp)
	}
	return false
}

func putFix[K comparable, V any](c int8, t **Node[K, V]) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return true
	}

	if s.b == -c {
		s.b = 0
		return false
	}

	if s.Children[(c+1)/2].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return false
}

func removeFix[K comparable, V any](c int8, t **Node[K, V]) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return false
	}

	if s.b == -c {
		s.b = 0
		return true
	}

	a := (c + 1) / 2
	if s.Children[a].b == 0 {
		s = rotate(c, s)
		s.b = -c
		*t = s
		return false
	}

	if s.Children[a].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return true
}

func singlerot[K comparable, V any](c int8, s *Node[K, V]) *Node[K, V] {
	s.b = 0
	s = rotate(c, s)
	s.b = 0
	return s
}

func doublerot[K comparable, V any](c int8, s *Node[K, V]) *Node[K, V] {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = rotate(-c, s.Children[a])
	p := rotate(c, s)

	switch {
	default:
		s.b = 0
		r.b = 0
	case p.b == c:
		s.b = -c
		r.b = 0
	case p.b == -c:
		s.b = 0
		r.b = c
	}

	p.b = 0
	return p
}

func rotate[K comparable, V any](c int8, s *Node[K, V]) *Node[K, V] {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = r.Children[a^1]
	if s.Children[a] != nil {
		s.Children[a].Parent = s
	}
	r.Children[a^1] = s
	r.Parent = s.Parent
	s.Parent = r
	return r
}

func (tree *Tree[K, V]) bottom(d int) *Node[K, V] {
	n := tree.Root
	if n == nil {
		return nil
	}

	for c := n.Children[d]; c != nil; c = n.Children[d] {
		n = c
	}
	return n
}

// Prev returns the previous element in an inorder
// walk of the AVL tree.
func (n *Node[K, V]) Prev() *Node[K, V] {
	return n.walk1(0)
}

// Next returns the next element in an inorder
// walk of the AVL tree.
func (n *Node[K, V]) Next() *Node[K, V] {
	return n.walk1(1)
}

func (n *Node[K, V]) walk1(a int) *Node[K, V] {
	if n == nil {
		return nil
	}

	if n.Children[a] != nil {
		n = n.Children[a]
		for n.Children[a^1] != nil {
			n = n.Children[a^1]
		}
		return n
	}

	p := n.Parent
	for p != nil && p.Children[a] == n {
		n = p
		p = p.Parent
	}
	return p
}

func output[K comparable, V any](node *Node[K, V], prefix string, isTail bool, str *string) {
	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Children[0], newPrefix, true, str)
	}
}
