// Copyright (c) 2017, Benjamin Scher Purcell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package avltree implements an AVL balanced binary tree.
//
// Structure is not thread safe.
package avltree

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/emirpasic/gods/utils"
	"github.com/emirpasic/gods/trees"
)

func assertTreeImplementation() {
	var _ trees.Tree = new(Tree)
}

var dbgLog = log.New(ioutil.Discard, "avltree: ", log.LstdFlags)

// Tree holds elements of the AVL tree.
type Tree struct {
	Root       *Node
	size       int
	comparator utils.Comparator
}

// Node is a single element within the tree
type Node struct {
	Key   interface{}
	Value interface{}
	c     [2]*Node
	p     *Node
	b     int8
}

// NewWith instantiates an AVL tree with the custom comparator.
func NewWith(comparator utils.Comparator) *Tree {
	return &Tree{comparator: comparator}
}

// NewWithIntComparator instantiates an AVL tree with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator() *Tree {
	return &Tree{comparator: utils.IntComparator}
}

// NewWithStringComparator instantiates an AVL tree with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator() *Tree {
	return &Tree{comparator: utils.StringComparator}
}

// Comparator returns the comparator function for the tree.
func (t *Tree) Comparator() utils.Comparator {
	return t.comparator
}

// New returns a new empty tree with the same comparator.
func (t *Tree) New() trees.Tree {
	return &Tree{comparator: t.comparator}
}

// Size returns the number of elements stored in the tree.
func (t *Tree) Size() int {
	return t.size
}

// Empty returns true if tree does not contain any nodes.
func (t *Tree) Empty() bool {
	return t.size == 0
}

// Clear removes all nodes from the tree.
func (t *Tree) Clear() {
	t.Root = nil
	t.size = 0
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree) Get(key interface{}) (value interface{}, found bool) {
	n := t.Root
	for n != nil {
		cmp := t.comparator(key, n.Key)
		switch {
		case cmp == 0:
			return n.Value, true
		case cmp < 0:
			n = n.c[0]
		case cmp > 0:
			n = n.c[1]
		}
	}
	return nil, false
}

// Floor Finds floor node of the input key, return the floor node or nil if no ceiling is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree is larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree) Floor(key interface{}) (floor *Node, found bool) {
	found = false
	n := t.Root
	for n != nil {
		c := t.comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			n = n.c[0]
		case c > 0:
			floor, found = n, true
			n = n.c[1]
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
func (t *Tree) Ceiling(key interface{}) (floor *Node, found bool) {
	found = false
	n := t.Root
	for n != nil {
		c := t.comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			floor, found = n, true
			n = n.c[0]
		case c > 0:
			n = n.c[1]
		}
	}
	if found {
		return
	}
	return nil, false
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree) Put(key interface{}, value interface{}) {
	var put func(*Node, **Node) bool
	put = func(p *Node, qp **Node) bool {
		q := *qp
		if q == nil {
			t.size++
			*qp = &Node{Key: key, Value: value, p: p}
			return true
		}

		c := t.comparator(key, q.Key)
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
		fix = put(q, &q.c[a])
		if fix {
			return putFix(int8(c), qp)
		}
		return false
	}

	put(nil, &t.Root)
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree) Remove(key interface{}) {
	var remove func(**Node) bool
	remove = func(qp **Node) bool {
		q := *qp
		if q == nil {
			return false
		}

		c := t.comparator(key, q.Key)
		if c == 0 {
			t.size--
			if q.c[1] == nil {
				if q.c[0] != nil {
					q.c[0].p = q.p
				}
				*qp = q.c[0]
				return true
			}
			fix := removeMin(&q.c[1], &q.Key, &q.Value)
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
		fix := remove(&q.c[a])
		if fix {
			return removeFix(int8(-c), qp)
		}
		return false
	}

	remove(&t.Root)
}

func removeMin(qp **Node, minKey *interface{}, minVal *interface{}) bool {
	q := *qp
	if q.c[0] == nil {
		*minKey = q.Key
		*minVal = q.Value
		if q.c[1] != nil {
			q.c[1].p = q.p
		}
		*qp = q.c[1]
		return true
	}
	fix := removeMin(&q.c[0], minKey, minVal)
	if fix {
		return removeFix(1, qp)
	}
	return false
}

func putFix(c int8, t **Node) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return true
	}

	if s.b == -c {
		s.b = 0
		return false
	}

	if s.c[(c+1)/2].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return false
}

func removeFix(c int8, t **Node) bool {
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
	if s.c[a].b == 0 {
		s = rotate(c, s)
		s.b = -c
		*t = s
		return false
	}

	if s.c[a].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return true
}

func singlerot(c int8, s *Node) *Node {
	dbgLog.Printf("singlerot: enter %p:%v %d\n", s, s, c)
	s.b = 0
	s = rotate(c, s)
	s.b = 0
	dbgLog.Printf("singlerot: exit %p:%v\n", s, s)
	return s
}

func doublerot(c int8, s *Node) *Node {
	dbgLog.Printf("doublerot: enter %p:%v %d\n", s, s, c)
	a := (c + 1) / 2
	r := s.c[a]
	s.c[a] = rotate(-c, s.c[a])
	p := rotate(c, s)
	if r.p != p || s.p != p {
		panic("doublerot: bad parents")
	}

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
	dbgLog.Printf("doublerot: exit %p:%v\n", s, s)
	return p
}

func rotate(c int8, s *Node) *Node {
	dbgLog.Printf("rotate: enter %p:%v %d\n", s, s, c)
	a := (c + 1) / 2
	r := s.c[a]
	s.c[a] = r.c[a^1]
	if s.c[a] != nil {
		s.c[a].p = s
	}
	r.c[a^1] = s
	r.p = s.p
	s.p = r
	dbgLog.Printf("rotate: exit %p:%v\n", r, r)
	return r
}

// Keys returns all keys in-order
func (t *Tree) Keys() []interface{} {
	keys := make([]interface{}, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

// Values returns all values in-order based on the key.
func (t *Tree) Values() []interface{} {
	values := make([]interface{}, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

// Left returns the minimum element of the AVL tree
// or nil if the tree is empty.
func (t *Tree) Left() *Node {
	return t.bottom(0)
}

// Right returns the maximum element of the AVL tree
// or nil if the tree is empty.
func (t *Tree) Right() *Node {
	return t.bottom(1)
}

// Min returns the minimum key value pair of the AVL tree
// or nils if the tree is empty.
func (t *Tree) Min() (interface{}, interface{}) {
	n := t.bottom(0)
	if n == nil {
		return nil, nil
	}
	return n.Key, n.Value
}

// Max returns the minimum key value pair of the AVL tree
// or nils if the tree is empty.
func (t *Tree) Max() (interface{}, interface{}) {
	n := t.bottom(1)
	if n == nil {
		return nil, nil
	}
	return n.Key, n.Value
}

func (t *Tree) bottom(d int) *Node {
	n := t.Root
	if n == nil {
		return nil
	}

	for c := n.c[d]; c != nil; c = n.c[d] {
		n = c
	}
	return n
}

// Prev returns the previous element in an inorder
// walk of the AVL tree.
func (n *Node) Prev() *Node {
	return n.walk1(0)
}

// Next returns the next element in an inorder
// walk of the AVL tree.
func (n *Node) Next() *Node {
	return n.walk1(1)
}

func (n *Node) walk1(a int) *Node {
	if n == nil {
		return nil
	}

	if n.c[a] != nil {
		n = n.c[a]
		for n.c[a^1] != nil {
			n = n.c[a^1]
		}
		return n
	}

	p := n.p
	for p != nil && p.c[a] == n {
		n = p
		p = p.p
	}
	return p
}

// String returns a string representation of container
func (t *Tree) String() string {
	str := "AVLTree\n"
	if !t.Empty() {
		output(t.Root, "", true, &str)
	}
	return str
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Key)
}

func output(node *Node, prefix string, isTail bool, str *string) {
	if node.c[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.c[0], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.c[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.c[1], newPrefix, true, str)
	}
}
