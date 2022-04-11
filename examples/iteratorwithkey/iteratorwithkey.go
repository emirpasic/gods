// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"strings"
)

// IteratorWithKeyExample to demonstrate basic usage of IteratorWithKey
func main() {
	m := treemap.NewWithIntComparator()
	m.Put(0, "a")
	m.Put(1, "b")
	m.Put(2, "c")
	it := m.Iterator()

	fmt.Print("\nForward iteration\n")
	for it.Next() {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]") // [0:a][1:b][2:c]
	}

	fmt.Print("\nForward iteration (again)\n")
	for it.Begin(); it.Next(); {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]") // [0:a][1:b][2:c]
	}

	fmt.Print("\nBackward iteration\n")
	for it.Prev() {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]") // [2:c][1:b][0:a]
	}

	fmt.Print("\nBackward iteration (again)\n")
	for it.End(); it.Prev(); {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]") // [2:c][1:b][0:a]
	}

	if it.First() {
		fmt.Print("\nFirst key: ", it.Key())     // First key: 0
		fmt.Print("\nFirst value: ", it.Value()) // First value: a
	}

	if it.Last() {
		fmt.Print("\nLast key: ", it.Key())     // Last key: 2
		fmt.Print("\nLast value: ", it.Value()) // Last value: c
	}

	// Seek key-value pair whose value starts with "b"
	seek := func(key interface{}, value interface{}) bool {
		return strings.HasSuffix(value.(string), "b")
	}

	it.Begin()
	for found := it.NextTo(seek); found; found = it.Next() {
		fmt.Print("\nNextTo key: ", it.Key())
		fmt.Print("\nNextTo value: ", it.Value())
	} /*
		NextTo key: 1
		NextTo value: "b"
		NextTo key: 2
		NextTo value: "c"
	*/

	it.End()
	for found := it.PrevTo(seek); found; found = it.Prev() {
		fmt.Print("\nNextTo key: ", it.Key())
		fmt.Print("\nNextTo value: ", it.Value())
	} /*
		NextTo key: 1
		NextTo value: "b"
		NextTo key: 0
		NextTo value: "a"
	*/
}
