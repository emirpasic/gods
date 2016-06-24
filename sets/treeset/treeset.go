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

// Implementation of an ordered set backed by a red-black tree.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

package treeset

import (
	"fmt"
	"github.com/emirpasic/gods/containers"
	"github.com/emirpasic/gods/sets"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
	"strings"
)

func assertInterfaceImplementation() {
	var _ sets.Set = (*Set)(nil)
	var _ containers.EnumerableWithIndex = (*Set)(nil)
	var _ containers.IteratorWithIndex = (*Iterator)(nil)
}

type Set struct {
	tree *rbt.Tree
}

var itemExists = struct{}{}

// Instantiates a new empty set with the custom comparator.
func NewWith(comparator utils.Comparator) *Set {
	return &Set{tree: rbt.NewWith(comparator)}
}

// Instantiates a new empty set with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator() *Set {
	return &Set{tree: rbt.NewWithIntComparator()}
}

// Instantiates a new empty set with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator() *Set {
	return &Set{tree: rbt.NewWithStringComparator()}
}

// Adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.tree.Put(item, itemExists)
	}
}

// Removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Check weather items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, contains := set.tree.Get(item); !contains {
			return false
		}
	}
	return true
}

// Returns true if set does not contain any elements.
func (set *Set) Empty() bool {
	return set.tree.Size() == 0
}

// Returns number of elements within the set.
func (set *Set) Size() int {
	return set.tree.Size()
}

// Clears all values in the set.
func (set *Set) Clear() {
	set.tree.Clear()
}

// Returns all items in the set.
func (set *Set) Values() []interface{} {
	return set.tree.Keys()
}

type Iterator struct {
	index    int
	iterator rbt.Iterator
}

// Returns a stateful iterator whose values can be fetched by an index.
func (set *Set) Iterator() Iterator {
	return Iterator{index: -1, iterator: set.tree.Iterator()}
}

// Moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator) Next() bool {
	iterator.index += 1
	return iterator.iterator.Next()
}

// Returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	return iterator.iterator.Key()
}

// Returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}

// Calls the given function once for each element, passing that element's index and value.
func (set *Set) Each(f func(index int, value interface{})) {
	iterator := set.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (set *Set) Map(f func(index int, value interface{}) interface{}) *Set {
	newSet := &Set{tree: rbt.NewWith(set.tree.Comparator)}
	iterator := set.Iterator()
	for iterator.Next() {
		newSet.Add(f(iterator.Index(), iterator.Value()))
	}
	return newSet
}

// Returns a new container containing all elements for which the given function returns a true value.
func (set *Set) Select(f func(index int, value interface{}) bool) *Set {
	newSet := &Set{tree: rbt.NewWith(set.tree.Comparator)}
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newSet.Add(iterator.Value())
		}
	}
	return newSet
}

// Passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (set *Set) Any(f func(index int, value interface{}) bool) bool {
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// Passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (set *Set) All(f func(index int, value interface{}) bool) bool {
	iterator := set.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (set *Set) Find(f func(index int, value interface{}) bool) (int, interface{}) {
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, nil
}

func (set *Set) String() string {
	str := "TreeSet\n"
	items := []string{}
	for _, v := range set.tree.Keys() {
		items = append(items, fmt.Sprintf("%v", v))
	}
	str += strings.Join(items, ", ")
	return str
}
