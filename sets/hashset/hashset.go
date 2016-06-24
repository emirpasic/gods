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
	var _ sets.Set = (*Set)(nil)
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

// Check if items (one or more) are present in the set.
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
