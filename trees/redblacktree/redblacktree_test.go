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

	tree.Put(1, "a")
	tree.Put(2, "b")
	tree.Clear()

	// Test Empty()
	if actualValue := tree.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

}

func BenchmarkRedBlackTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := NewWithIntComparator()
		for n := 0; n < 1000; n++ {
			tree.Put(n, n)
		}
		for n := 0; n < 1000; n++ {
			tree.Remove(n)
		}
	}
}
