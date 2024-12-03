// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package priorityqueue_test

import (
	"cmp"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/emirpasic/gods/v2/queues/priorityqueue"
)

type Element struct {
	priority int
	name     string
}

func (element Element) String() string {
	return fmt.Sprintf("{%v %v}", element.priority, element.name)
}

// Comparator function (sort by priority value in descending order)
func byPriority(a, b Element) int {
	return -cmp.Compare(a.priority, b.priority) // Note "-" for descending order
}

func TestBinaryQueueEnqueue(t *testing.T) {
	queue := priorityqueue.NewWith[Element](byPriority)

	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	a := Element{name: "a", priority: 1}
	c := Element{name: "c", priority: 3}
	b := Element{name: "b", priority: 2}

	queue.Enqueue(a)
	queue.Enqueue(c)
	queue.Enqueue(b)

	it := queue.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value.name, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value.name, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value.name, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := index, count-1; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}

	if actualValue := queue.Values(); actualValue[0].name != "c" || actualValue[1].name != "b" || actualValue[2].name != "a" {
		t.Errorf("Got %v expected %v", actualValue, `[{3 c} {2 b} {1 a}]`)
	}
}

func TestBinaryQueueEnqueueBulk(t *testing.T) {
	queue := priorityqueue.New[int]()

	queue.Enqueue(15)
	queue.Enqueue(20)
	queue.Enqueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)

	if actualValue, ok := queue.Dequeue(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 15 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 15)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 20 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 20)
	}

	queue.Clear()
	if actualValue := queue.Empty(); !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestBinaryQueueDequeue(t *testing.T) {
	queue := priorityqueue.New[int]()

	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)
	queue.Dequeue() // removes 1

	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := queue.Values(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}
}

func TestBinaryQueueRandom(t *testing.T) {
	queue := priorityqueue.New[int]()

	rand.Seed(3)
	for i := 0; i < 10000; i++ {
		r := int(rand.Int31n(30))
		queue.Enqueue(r)
	}

	prev, _ := queue.Dequeue()
	for !queue.Empty() {
		curr, _ := queue.Dequeue()
		if prev > curr {
			t.Errorf("Queue property invalidated. prev: %v current: %v", prev, curr)
		}
		prev = curr
	}
}

func TestBinaryQueueIteratorOnEmpty(t *testing.T) {
	queue := priorityqueue.New[int]()
	it := queue.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty queue")
	}
}

func TestBinaryQueueIteratorNext(t *testing.T) {
	queue := priorityqueue.New[int]()
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)

	it := queue.Iterator()
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

func TestBinaryQueueIteratorPrev(t *testing.T) {
	queue := priorityqueue.New[int]()
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)

	it := queue.Iterator()
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

func TestBinaryQueueIteratorBegin(t *testing.T) {
	queue := priorityqueue.New[int]()
	it := queue.Iterator()
	it.Begin()
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(1)
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != 1 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, 1)
	}
}

func TestBinaryQueueIteratorEnd(t *testing.T) {
	queue := priorityqueue.New[int]()
	it := queue.Iterator()

	if index := it.Index(); index != -1 {
		t.Errorf("Got %v expected %v", index, -1)
	}

	it.End()
	if index := it.Index(); index != 0 {
		t.Errorf("Got %v expected %v", index, 0)
	}

	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)
	it.End()
	if index := it.Index(); index != queue.Size() {
		t.Errorf("Got %v expected %v", index, queue.Size())
	}

	it.Prev()
	if index, value := it.Index(), it.Value(); index != queue.Size()-1 || value != 3 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, queue.Size()-1, 3)
	}
}

func TestBinaryQueueIteratorFirst(t *testing.T) {
	queue := priorityqueue.New[int]()
	it := queue.Iterator()
	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	queue.Enqueue(3) // [3]
	queue.Enqueue(2) // [2,3]
	queue.Enqueue(1) // [1,3,2](2 swapped with 1, hence last)
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != 1 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, 1)
	}
}

func TestBinaryQueueIteratorLast(t *testing.T) {
	tree := priorityqueue.New[int]()
	it := tree.Iterator()
	if actualValue, expectedValue := it.Last(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	tree.Enqueue(2)
	tree.Enqueue(3)
	tree.Enqueue(1)
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 2 || value != 3 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 2, 3)
	}
}

func TestBinaryQueueIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// NextTo (empty)
	{
		tree := priorityqueue.New[string]()
		it := tree.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// NextTo (not found)
	{
		tree := priorityqueue.New[string]()
		tree.Enqueue("xx")
		tree.Enqueue("yy")
		it := tree.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// NextTo (found)
	{
		tree := priorityqueue.New[string]()
		tree.Enqueue("aa")
		tree.Enqueue("bb")
		tree.Enqueue("cc")
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

func TestBinaryQueueIteratorPrevTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// PrevTo (empty)
	{
		tree := priorityqueue.New[string]()
		it := tree.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// PrevTo (not found)
	{
		tree := priorityqueue.New[string]()
		tree.Enqueue("xx")
		tree.Enqueue("yy")
		it := tree.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// PrevTo (found)
	{
		tree := priorityqueue.New[string]()
		tree.Enqueue("aa")
		tree.Enqueue("bb")
		tree.Enqueue("cc")
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

func TestBinaryQueueSerialization(t *testing.T) {
	queue := priorityqueue.New[string]()

	queue.Enqueue("c")
	queue.Enqueue("b")
	queue.Enqueue("a")

	var err error
	assert := func() {
		if actualValue := queue.Values(); actualValue[0] != "a" || actualValue[1] != "b" || actualValue[2] != "c" {
			t.Errorf("Got %v expected %v", actualValue, "[1,3,2]")
		}
		if actualValue := queue.Size(); actualValue != 3 {
			t.Errorf("Got %v expected %v", actualValue, 3)
		}
		if actualValue, ok := queue.Peek(); actualValue != "a" || !ok {
			t.Errorf("Got %v expected %v", actualValue, "a")
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	bytes, err := queue.ToJSON()
	assert()

	err = queue.FromJSON(bytes)
	assert()

	bytes, err = json.Marshal([]interface{}{"a", "b", "c", queue})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	err = json.Unmarshal([]byte(`["a","b","c"]`), &queue)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	assert()
}

func TestBTreeString(t *testing.T) {
	c := priorityqueue.New[int]()
	c.Enqueue(1)
	if !strings.HasPrefix(c.String(), "PriorityQueue") {
		t.Errorf("String should start with container name")
	}
}

func benchmarkEnqueue(b *testing.B, queue *priorityqueue.Queue[Element], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Enqueue(Element{})
		}
	}
}

func benchmarkDequeue(b *testing.B, queue *priorityqueue.Queue[Element], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Dequeue()
		}
	}
}

func BenchmarkBinaryQueueDequeue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := priorityqueue.NewWith[Element](byPriority)
	for n := 0; n < size; n++ {
		queue.Enqueue(Element{})
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkBinaryQueueDequeue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := priorityqueue.NewWith[Element](byPriority)
	for n := 0; n < size; n++ {
		queue.Enqueue(Element{})
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkBinaryQueueDequeue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := priorityqueue.NewWith[Element](byPriority)
	for n := 0; n < size; n++ {
		queue.Enqueue(Element{})
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkBinaryQueueDequeue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := priorityqueue.NewWith[Element](byPriority)
	for n := 0; n < size; n++ {
		queue.Enqueue(Element{})
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkBinaryQueueEnqueue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := priorityqueue.NewWith(byPriority)
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkBinaryQueueEnqueue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := priorityqueue.NewWith[Element](byPriority)
	for n := 0; n < size; n++ {
		queue.Enqueue(Element{})
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkBinaryQueueEnqueue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := priorityqueue.NewWith[Element](byPriority)
	for n := 0; n < size; n++ {
		queue.Enqueue(Element{})
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkBinaryQueueEnqueue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := priorityqueue.NewWith[Element](byPriority)
	for n := 0; n < size; n++ {
		queue.Enqueue(Element{})
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}
