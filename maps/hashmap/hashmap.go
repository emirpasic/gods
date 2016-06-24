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
	var _ maps.Map = (*Map)(nil)
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
