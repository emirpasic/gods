// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktree

import (
	"fmt"
	"testing"
)

func TestRedBlackTreePut(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(8, "h")
	tree.Put(9, "i")
	tree.Put(12, "l")
	tree.Put(10, "j")
	tree.Put(17, "q")
	tree.Put(15, "o")
	tree.Put(19, "s")
	tree.Put(1, "a") //overwrite
	//│       ┌── 19
	//│       │   └──(17)
	//│   ┌── 15
	//│   │   │   ┌── 12
	//│   │   └──(10)
	//│   │       └── 9
	//└── 8
	//    │       ┌── 7
	//	  │   ┌── 6
	//	  │   │   └── 5
	//	  └──(4)
	//	      │   ┌── 3
	//		  └── 2
	//		      └── 1

	if actualValue := tree.Size(); actualValue != 14 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	if actualValue, expectedValue := fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d", tree.Keys()...), "1234567891012151719"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s", tree.Values()...), "abcdefghijloqs"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, "h", true},
		{9, "i", true},
		{10, "j", true},
		{12, "l", true},
		{15, "o", true},
		{17, "q", true},
		{19, "s", true},
		{20, nil, false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := tree.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	tree.Validate()
	tree.Clear()
	tree.Validate()
}

func TestRedBlackTreeRemove(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	tree.Remove(5)
	tree.Validate()
	tree.Remove(6)
	tree.Validate()
	tree.Remove(7)
	tree.Validate()
	tree.Remove(8)
	tree.Validate()
	tree.Remove(5)
	tree.Validate()

	if actualValue, expectedValue := fmt.Sprintf("%d%d%d%d", tree.Keys()...), "1234"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", tree.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", tree.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
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
		actualValue, actualFound := tree.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	tree.Remove(1)
	tree.Remove(4)
	tree.Remove(2)
	tree.Remove(3)
	tree.Remove(2)
	tree.Remove(2)

	if actualValue, expectedValue := fmt.Sprintf("%s", tree.Keys()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", tree.Values()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if empty, size := tree.Empty(), tree.Size(); empty != true || size != -0 {
		t.Errorf("Got %v expected %v", empty, true)
	}
}

func TestRedBlackTreeLeftAndRight(t *testing.T) {
	tree := NewWithIntComparator()

	if actualValue := tree.Left(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := tree.Right(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

	tree.Put(1, "a")
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x") // overwrite
	tree.Put(2, "b")

	if actualValue, expectedValue := fmt.Sprintf("%d", tree.Left().Key), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", tree.Left().Value), "x"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, expectedValue := fmt.Sprintf("%d", tree.Right().Key), "7"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", tree.Right().Value), "g"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeCeilingAndFloor(t *testing.T) {
	tree := NewWithIntComparator()

	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
	if node, found := tree.Ceiling(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(8, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	//
	//│       ┌── 8
	//│   ┌── 6
	//│   │   └── 5
	//└── 4
	//    │   ┌── 3
	//    └── 2
	//        └── 1

	if node, found := tree.Floor(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	if node, found := tree.Floor(7); node.Key != 6 || !found {
		t.Errorf("Got %v expected %v", node.Key, 6)
	}

	if node, found := tree.Ceiling(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}

	if node, found := tree.Ceiling(7); node.Key != 8 || !found {
		t.Errorf("Got %v expected %v", node.Key, 8)
	}

	if node, found := tree.Ceiling(9); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
}

func benchmarkGet(b *testing.B, tree *Tree, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, tree *Tree, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, tree *Tree, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Remove(n)
		}
	}
}

func BenchmarkRedBlackTreeGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreePut100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator()
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator()
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator()
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator()
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}
