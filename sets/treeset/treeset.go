// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treeset implements a tree backed by a red-black tree.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package treeset

import (
	"cmp"
	"fmt"
	"reflect"
	"strings"

	"github.com/emirpasic/gods/v2/sets"
	rbt "github.com/emirpasic/gods/v2/trees/redblacktree"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert Set implementation
var _ sets.Set[int] = (*Set[int])(nil)

// Set holds elements in a red-black tree
type Set[T comparable] struct {
	tree *rbt.Tree[T, struct{}]
}

var itemExists = struct{}{}

func New[T cmp.Ordered](values ...T) *Set[T] {
	return NewWith[T](cmp.Compare[T], values...)
}

// NewWith instantiates a new empty set with the custom comparator.
func NewWith[T comparable](comparator utils.Comparator[T], values ...T) *Set[T] {
	set := &Set[T]{tree: rbt.NewWith[T, struct{}](comparator)}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.tree.Put(item, itemExists)
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[T]) Remove(items ...T) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Contains checks weather items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.tree.Get(item); !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[T]) Empty() bool {
	return set.tree.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[T]) Size() int {
	return set.tree.Size()
}

// Clear clears all values in the set.
func (set *Set[T]) Clear() {
	set.tree.Clear()
}

// Values returns all items in the set.
func (set *Set[T]) Values() []T {
	return set.tree.Keys()
}

// String returns a string representation of container
func (set *Set[T]) String() string {
	str := "TreeSet\n"
	items := []string{}
	for _, v := range set.tree.Keys() {
		items = append(items, fmt.Sprintf("%v", v))
	}
	str += strings.Join(items, ", ")
	return str
}

// Intersection returns the intersection between two sets.
// The new set consists of all elements that are both in "set" and "another".
// The two sets should have the same comparators, otherwise the result is empty set.
// Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (set *Set[T]) Intersection(another *Set[T]) *Set[T] {
	result := NewWith(set.tree.Comparator)

	setComparator := reflect.ValueOf(set.tree.Comparator)
	anotherComparator := reflect.ValueOf(another.tree.Comparator)
	if setComparator.Pointer() != anotherComparator.Pointer() {
		return result
	}

	// Iterate over smaller set (optimization)
	if set.Size() <= another.Size() {
		for it := set.Iterator(); it.Next(); {
			if another.Contains(it.Value()) {
				result.Add(it.Value())
			}
		}
	} else {
		for it := another.Iterator(); it.Next(); {
			if set.Contains(it.Value()) {
				result.Add(it.Value())
			}
		}
	}

	return result
}

// Union returns the union of two sets.
// The new set consists of all elements that are in "set" or "another" (possibly both).
// The two sets should have the same comparators, otherwise the result is empty set.
// Ref: https://en.wikipedia.org/wiki/Union_(set_theory)
func (set *Set[T]) Union(another *Set[T]) *Set[T] {
	result := NewWith(set.tree.Comparator)

	setComparator := reflect.ValueOf(set.tree.Comparator)
	anotherComparator := reflect.ValueOf(another.tree.Comparator)
	if setComparator.Pointer() != anotherComparator.Pointer() {
		return result
	}

	for it := set.Iterator(); it.Next(); {
		result.Add(it.Value())
	}
	for it := another.Iterator(); it.Next(); {
		result.Add(it.Value())
	}

	return result
}

// Difference returns the difference between two sets.
// The two sets should have the same comparators, otherwise the result is empty set.
// The new set consists of all elements that are in "set" but not in "another".
// Ref: https://proofwiki.org/wiki/Definition:Set_Difference
func (set *Set[T]) Difference(another *Set[T]) *Set[T] {
	result := NewWith(set.tree.Comparator)

	setComparator := reflect.ValueOf(set.tree.Comparator)
	anotherComparator := reflect.ValueOf(another.tree.Comparator)
	if setComparator.Pointer() != anotherComparator.Pointer() {
		return result
	}

	for it := set.Iterator(); it.Next(); {
		if !another.Contains(it.Value()) {
			result.Add(it.Value())
		}
	}

	return result
}
