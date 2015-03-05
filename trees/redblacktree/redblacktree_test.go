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

package redblacktree

import (
	"fmt"
	"testing"
)

func TestRedBlackTree(t *testing.T) {

	tree := NewWithIntComparator()

	// insertions
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	// Test Size()
	if actualValue := tree.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	// test Keys()
	if actualValue, expactedValue := fmt.Sprintf("%d%d%d%d%d%d%d", tree.Keys()...), "1234567"; actualValue != expactedValue {
		t.Errorf("Got %v expected %v", actualValue, expactedValue)
	}

	// test Values()
	if actualValue, expactedValue := fmt.Sprintf("%s%s%s%s%s%s%s", tree.Values()...), "abcdefg"; actualValue != expactedValue {
		t.Errorf("Got %v expected %v", actualValue, expactedValue)
	}

	// key,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, nil, false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := tree.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	// removals
	tree.Remove(5)
	tree.Remove(6)
	tree.Remove(7)
	tree.Remove(8)
	tree.Remove(5)

	// Test Keys()
	if actualValue, expactedValue := fmt.Sprintf("%d%d%d%d", tree.Keys()...), "1234"; actualValue != expactedValue {
		t.Errorf("Got %v expected %v", actualValue, expactedValue)
	}

	// test Values()
	if actualValue, expactedValue := fmt.Sprintf("%s%s%s%s", tree.Values()...), "abcd"; actualValue != expactedValue {
		t.Errorf("Got %v expected %v", actualValue, expactedValue)
	}

	// Test Size()
	if actualValue := tree.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	tests2 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, nil, false},
		{6, nil, false},
		{7, nil, false},
		{8, nil, false},
	}

	for _, test := range tests2 {
		// retrievals
		actualValue, actualFound := tree.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	// removals
	tree.Remove(1)
	tree.Remove(4)
	tree.Remove(2)
	tree.Remove(3)
	tree.Remove(2)
	tree.Remove(2)

	// Test Keys()
	if actualValue, expactedValue := fmt.Sprintf("", tree.Keys()...), ""; actualValue != expactedValue {
		t.Errorf("Got %v expected %v", actualValue, expactedValue)
	}

	// test Values()
	if actualValue, expactedValue := fmt.Sprintf("", tree.Values()...), ""; actualValue != expactedValue {
		t.Errorf("Got %v expected %v", actualValue, expactedValue)
	}

	// Test Size()
	if actualValue := tree.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	// Test Empty()
	if actualValue := tree.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

}
