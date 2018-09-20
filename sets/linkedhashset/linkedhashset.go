// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedhashset implements a set that preserves insertion-order and is backed a hash tables to store values and doubly-linked list to store ordering.
//
// Note that insertion-order is not affected if an element is re-inserted into the set.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package linkedhashset

import (
	"fmt"
	"github.com/emirpasic/gods/lists/doublylinkedlist"
	"github.com/emirpasic/gods/sets"
	"strings"
)

func assertSetImplementation() {
	var _ sets.Set = (*Set)(nil)
}

// Set holds elements in go's native map
type Set struct {
	items map[interface{}]struct{}
	list  *doublylinkedlist.List
}

var itemExists = struct{}{}

// New instantiates a new empty set
func New() *Set {
	return &Set{
		items: make(map[interface{}]struct{}),
		list:  doublylinkedlist.New(),
	}
}

// Add adds the items (one or more) to the set.
// Note that insertion-order is not affected if an element is re-inserted into the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {

		if _, contains := set.items[item]; contains {
			continue
		}

		set.items[item] = itemExists
		set.list.Append(item)
	}
}

// Remove removes the items (one or more) from the set.
// Slow operation, worst-case O(n^2).
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {

		if _, contains := set.items[item]; !contains {
			continue
		}

		delete(set.items, item)
		index := set.list.IndexOf(item)
		set.list.Remove(index)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return set.list.Size()
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.items = make(map[interface{}]struct{})
	set.list.Clear()
}

// Values returns all items in the set.
func (set *Set) Values() []interface{} {
	values := make([]interface{}, set.Size())
	it := set.Iterator()
	for it.Next() {
		values[it.Index()] = it.Value()
	}
	return values
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "LinkedHashSet\n"
	items := []string{}
	it := set.Iterator()
	for it.Next() {
		items = append(items, fmt.Sprintf("%v", it.Value()))
	}
	str += strings.Join(items, ", ")
	return str
}
