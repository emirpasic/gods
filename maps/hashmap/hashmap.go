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

// Implementation of unorder map backed by a hash table.
// Elements are unordered in the map.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Associative_array

package hashmap

import (
	"fmt"
	"github.com/emirpasic/gods/maps"
)

func assertInterfaceImplementation() {
	var _ maps.Interface = (*Map)(nil)
}

type Map struct {
	m map[interface{}]interface{}
}

// Instantiates a hash map.
func New() *Map {
	return &Map{m: make(map[interface{}]interface{})}
}

// Inserts element into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	m.m[key] = value
}

// Searches the elemnt in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	value, found = m.m[key]
	return
}

// Remove the element from the map by key.
func (m *Map) Remove(key interface{}) {
	delete(m.m, key)
}

// Returns true if map does not contain any elements
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Returns number of elements in the map.
func (m *Map) Size() int {
	return len(m.m)
}

// Returns all keys (random order).
func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	count := 0
	for key, _ := range m.m {
		keys[count] = key
		count += 1
	}
	return keys
}

// Returns all values (random order).
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count += 1
	}
	return values
}

// Removes all elements from the map.
func (m *Map) Clear() {
	m.m = make(map[interface{}]interface{})
}

func (m *Map) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}
