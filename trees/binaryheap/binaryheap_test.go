// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import (
	"encoding/json"
	"math/rand"
	"slices"
	"strings"
	"testing"
)

func TestBinaryHeapPush(t *testing.T) {
	heap := New[int]()

	if actualValue := heap.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	heap.Push(3)
	heap.Push(2)
	heap.Push(1)

	if actualValue, expectedValue := heap.Values(), []int{1, 2, 3}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := heap.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := heap.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := heap.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestBinaryHeapPushBulk(t *testing.T) {
	heap := New[int]()

	heap.Push(15, 20, 3, 1, 2)

	if actualValue, expectedValue := heap.Values(), []int{1, 2, 3, 15, 20}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, ok := heap.Pop(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestBinaryHeapPop(t *testing.T) {
	heap := New[int]()

	if actualValue := heap.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	heap.Push(3)
	heap.Push(2)
	heap.Push(1)
	heap.Pop()

	if actualValue, ok := heap.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := heap.Pop(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := heap.Pop(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := heap.Pop(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := heap.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := heap.Values(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}
}

func TestBinaryHeapRandom(t *testing.T) {
	heap := New[int]()

	rand.Seed(3)
	for i := 0; i < 10000; i++ {
		r := int(rand.Int31n(30))
		heap.Push(r)
	}

	prev, _ := heap.Pop()
	for !heap.Empty() {
		curr, _ := heap.Pop()
		if prev > curr {
			t.Errorf("Heap property invalidated. prev: %v current: %v", prev, curr)
		}
		prev = curr
	}
}

func TestBinaryHeapIteratorOnEmpty(t *testing.T) {
	heap := New[int]()
	it := heap.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty heap")
	}
}

func TestBinaryHeapIteratorNext(t *testing.T) {
	heap := New[int]()
	heap.Push(3)
	heap.Push(2)
	heap.Push(1)

	it := heap.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := index, count-1; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestBinaryHeapIteratorPrev(t *testing.T) {
	heap := New[int]()
	heap.Push(3)
	heap.Push(2)
	heap.Push(1)

	it := heap.Iterator()
	for it.Next() {
	}
	count := 0
	for it.Prev() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := index, 3-count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestBinaryHeapIteratorBegin(t *testing.T) {
	heap := New[int]()
	it := heap.Iterator()
	it.Begin()
	heap.Push(2)
	heap.Push(3)
	heap.Push(1)
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != 1 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, 1)
	}
}

func TestBinaryHeapIteratorEnd(t *testing.T) {
	heap := New[int]()
	it := heap.Iterator()

	if index := it.Index(); index != -1 {
		t.Errorf("Got %v expected %v", index, -1)
	}

	it.End()
	if index := it.Index(); index != 0 {
		t.Errorf("Got %v expected %v", index, 0)
	}

	heap.Push(3)
	heap.Push(2)
	heap.Push(1)
	it.End()
	if index := it.Index(); index != heap.Size() {
		t.Errorf("Got %v expected %v", index, heap.Size())
	}

	it.Prev()
	if index, value := it.Index(), it.Value(); index != heap.Size()-1 || value != 3 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, heap.Size()-1, 3)
	}
}

func TestBinaryHeapIteratorFirst(t *testing.T) {
	heap := New[int]()
	it := heap.Iterator()
	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	heap.Push(3)
	heap.Push(2)
	heap.Push(1)
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != 1 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, 1)
	}
}

func TestBinaryHeapIteratorLast(t *testing.T) {
	tree := New[int]()
	it := tree.Iterator()
	if actualValue, expectedValue := it.Last(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	tree.Push(2)
	tree.Push(3)
	tree.Push(1)
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 2 || value != 3 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 2, 3)
	}
}

func TestBinaryHeapIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// NextTo (empty)
	{
		tree := New[string]()
		it := tree.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// NextTo (not found)
	{
		tree := New[string]()
		tree.Push("xx")
		tree.Push("yy")
		it := tree.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// NextTo (found)
	{
		tree := New[string]()
		tree.Push("aa")
		tree.Push("bb")
		tree.Push("cc")
		it := tree.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
		if index, value := it.Index(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Next() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Index(), it.Value(); index != 2 || value != "cc" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "cc")
		}
		if it.Next() {
			t.Errorf("Should not go past last element")
		}
	}
}

func TestBinaryHeapIteratorPrevTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// PrevTo (empty)
	{
		tree := New[string]()
		it := tree.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// PrevTo (not found)
	{
		tree := New[string]()
		tree.Push("xx")
		tree.Push("yy")
		it := tree.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// PrevTo (found)
	{
		tree := New[string]()
		tree.Push("aa")
		tree.Push("bb")
		tree.Push("cc")
		it := tree.Iterator()
		it.End()
		if !it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
		if index, value := it.Index(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Prev() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Index(), it.Value(); index != 0 || value != "aa" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "aa")
		}
		if it.Prev() {
			t.Errorf("Should not go before first element")
		}
	}
}

func TestBinaryHeapSerialization(t *testing.T) {
	heap := New[string]()

	heap.Push("c")
	heap.Push("b")
	heap.Push("a")

	var err error
	assert := func() {
		if actualValue, expectedValue := heap.Values(), []string{"a", "b", "c"}; !slices.Equal(actualValue, expectedValue) {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue := heap.Size(); actualValue != 3 {
			t.Errorf("Got %v expected %v", actualValue, 3)
		}
		if actualValue, ok := heap.Peek(); actualValue != "a" || !ok {
			t.Errorf("Got %v expected %v", actualValue, "a")
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	bytes, err := heap.ToJSON()
	assert()

	err = heap.FromJSON(bytes)
	assert()

	bytes, err = json.Marshal([]interface{}{"a", "b", "c", heap})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	intHeap := New[int]()
	err = json.Unmarshal([]byte(`[1,2,3]`), &intHeap)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	if actualValue, expectedValue := intHeap.Values(), []int{1, 2, 3}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestBTreeString(t *testing.T) {
	c := New[int]()
	c.Push(1)
	if !strings.HasPrefix(c.String(), "BinaryHeap") {
		t.Errorf("String should start with container name")
	}
}

func benchmarkPush(b *testing.B, heap *Heap[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			heap.Push(n)
		}
	}
}

func benchmarkPop(b *testing.B, heap *Heap[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			heap.Pop()
		}
	}
}

func BenchmarkBinaryHeapPop100(b *testing.B) {
	b.StopTimer()
	size := 100
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, heap, size)
}

func BenchmarkBinaryHeapPop1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, heap, size)
}

func BenchmarkBinaryHeapPop10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, heap, size)
}

func BenchmarkBinaryHeapPop100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, heap, size)
}

func BenchmarkBinaryHeapPush100(b *testing.B) {
	b.StopTimer()
	size := 100
	heap := New[int]()
	b.StartTimer()
	benchmarkPush(b, heap, size)
}

func BenchmarkBinaryHeapPush1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, heap, size)
}

func BenchmarkBinaryHeapPush10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, heap, size)
}

func BenchmarkBinaryHeapPush100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, heap, size)
}
