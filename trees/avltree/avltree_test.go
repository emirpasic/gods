// Copyright (c) 2017, Benjamin Scher Purcell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package avltree

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

var rng *rand.Rand

func TestMain(m *testing.M) {
	seed := time.Now().UTC().UnixNano()
	rng = rand.New(rand.NewSource(seed))
	os.Exit(m.Run())
}

func TestAVLPut(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	if actualValue := tree.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	if actualValue, expectedValue := fmt.Sprintf("%d%d%d%d%d%d%d", tree.Keys()...), "1234567"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s%s%s%s", tree.Values()...), "abcdefg"; actualValue != expectedValue {
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
		{8, nil, false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := tree.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}
}

func TestAVLRemove(t *testing.T) {
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
	tree.Remove(6)
	tree.Remove(7)
	tree.Remove(8)
	tree.Remove(5)

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

func TestAVLLeftAndRight(t *testing.T) {
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

func TestAVLCeilingAndFloor(t *testing.T) {
	tree := NewWithIntComparator()

	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
	if node, found := tree.Ceiling(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")

	if node, found := tree.Floor(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	if node, found := tree.Ceiling(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Ceiling(8); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
}

func TestAVLIteratorNextOnEmpty(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestAVLIteratorPrevOnEmpty(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestAVLIterator1Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLIterator1Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLIterator2Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLIterator2Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLIterator3Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(1, "a")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLIterator3Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(1, "a")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLIterator4Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		value := it.Value()
		switch value {
		case count:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLIterator4Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := tree.Size()
	for it.Next() {
	}
	for it.Prev() {
		value := it.Value()
		switch value {
		case count:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		count--
	}
	if actualValue, expectedValue := count, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLIteratorBegin(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Begin()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	for it.Next() {
	}

	it.Begin()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Next()
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestAVLIteratorEnd(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.Iterator()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.End()
	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it.End()
	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Prev()
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestAVLIteratorFirst(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestAVLIteratorLast(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func newRandomIntTree(size, randMax int) *Tree {
	tree := NewWithIntComparator()
	for i := 0; i < size; i++ {
		tree.Put(rng.Intn(randMax), nil)
	}
	return tree
}

func (tree *Tree) benchmarkGet(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(n)
		}
	}
}

func (tree *Tree) benchmarkGetRandom(b *testing.B, size, randMax int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(rng.Intn(randMax))
		}
	}
}

func (tree *Tree) benchmarkPut(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(n, nil)
		}
		tree.Empty()
	}
}

func (tree *Tree) benchmarkPutRandom(b *testing.B, size, randMax int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(rng.Intn(randMax), nil)
		}
		tree.Empty()
	}
}

func (tree *Tree) benchmarkPutAndRemove(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(n, nil)
		}
		for n := 0; n < size; n++ {
			tree.Remove(n)
		}
		tree.Empty()
	}
}

func (tree *Tree) benchmarkPutAndRemoveRandom(b *testing.B, size int, randMax int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(rng.Intn(randMax), nil)
		}
		for n := 0; n < size; n++ {
			tree.Remove(rng.Intn(randMax))
		}
		tree.Empty()
	}
}

func BenchmarkAVLGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkGet(b, size)
}

func BenchmarkAVLGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkGet(b, size)
}

func BenchmarkAVLGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkGet(b, size)
}

func BenchmarkAVLGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkGet(b, size)
}

func BenchmarkAVLGetRandom100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkGetRandom(b, size, size*3)
}

func BenchmarkAVLGetRandom1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkGetRandom(b, size, size*3)
}

func BenchmarkAVLGetRandom10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkGetRandom(b, size, size*3)
}

func BenchmarkAVLGetRandom100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkGetRandom(b, size, size*3)
}

func BenchmarkAVLPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkPut(b, size)
}

func BenchmarkAVLPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkPut(b, size)
}

func BenchmarkAVLPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkPut(b, size)
}

func BenchmarkAVLPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkPut(b, size)
}

func BenchmarkAVLPutRandom100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkPutRandom(b, size, size*3)
}

func BenchmarkAVLPutRandom1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkPutRandom(b, size, size*3)
}

func BenchmarkAVLPutRandom10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkPutRandom(b, size, size*3)
}

func BenchmarkAVLPutRandom100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkPutRandom(b, size, size*3)
}

func BenchmarkAVLPutAndRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkPutAndRemove(b, size)
}

func BenchmarkAVLPutAndRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkPutAndRemove(b, size)
}

func BenchmarkAVLPutAndRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkPutAndRemove(b, size)
}

func BenchmarkAVLPutAndRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n, nil)
	}
	b.StartTimer()
	tree.benchmarkPutAndRemove(b, size)
}

func BenchmarkAVLPutAndRemoveRandom100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkPutAndRemoveRandom(b, size, size*3)
}

func BenchmarkAVLPutAndRemoveRandom1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkPutAndRemoveRandom(b, size, size*3)
}

func BenchmarkAVLPutAndRemoveRandom10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkPutAndRemoveRandom(b, size, size*3)
}

func BenchmarkAVLPutAndRemoveRandom100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := newRandomIntTree(size, size*3)
	b.StartTimer()
	tree.benchmarkPutAndRemoveRandom(b, size, size*3)
}
