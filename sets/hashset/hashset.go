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

// Implementation of set backed by a hash table.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

package hashset

import (
	"fmt"
	"github.com/emirpasic/gods/sets"
	"strings"
)

func assertInterfaceImplementation() {
	var _ sets.Interface = (*Set)(nil)
}

type Set struct {
	items map[interface{}]struct{}
}

var itemExists = struct{}{}

// Instantiates a new empty set
func New() *Set {
	return &Set{items: make(map[interface{}]struct{})}
}

// Adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

// Removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Check wether items (one or more) are present in the set.
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

// Returns true if set does not contain any elements.
func (set *Set) Empty() bool {
	return set.Size() == 0
}

// Returns number of elements within the set.
func (set *Set) Size() int {
	return len(set.items)
}

// Clears all values in the set.
func (set *Set) Clear() {
	set.items = make(map[interface{}]struct{})
}

// Returns all items in the set.
func (set *Set) Values() []interface{} {
	values := make([]interface{}, set.Size())
	count := 0
	for item, _ := range set.items {
		values[count] = item
		count += 1
	}
	return values
}

func (set *Set) String() string {
	str := "HashSet\n"
	items := []string{}
	for k, _ := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}
