// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package circularbuffer_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/emirpasic/gods/v2/queues/circularbuffer"
	"github.com/emirpasic/gods/v2/testutils"
)

func TestQueueEnqueue(t *testing.T) {
	queue := circularbuffer.New[int](3)
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if actualValue := queue.Values(); actualValue[0] != 1 || actualValue[1] != 2 || actualValue[2] != 3 {
		t.Errorf("Got %v expected %v", actualValue, "[1,2,3]")
	}
	if actualValue := queue.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := queue.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := queue.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestQueuePeek(t *testing.T) {
	queue := circularbuffer.New[int](3)
	if actualValue, ok := queue.Peek(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if actualValue, ok := queue.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestQueueDequeue(t *testing.T) {
	assert := func(actualValue interface{}, expectedValue interface{}) {
		if actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}

	queue := circularbuffer.New[int](3)
	assert(queue.Empty(), true)
	assert(queue.Empty(), true)
	assert(queue.Full(), false)
	assert(queue.Size(), 0)
	queue.Enqueue(1)
	assert(queue.Size(), 1)
	queue.Enqueue(2)
	assert(queue.Size(), 2)

	queue.Enqueue(3)
	assert(queue.Size(), 3)
	assert(queue.Empty(), false)
	assert(queue.Full(), true)

	queue.Dequeue()
	assert(queue.Size(), 2)

	if actualValue, ok := queue.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	assert(queue.Size(), 2)

	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	assert(queue.Size(), 1)

	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	assert(queue.Size(), 0)
	assert(queue.Empty(), true)
	assert(queue.Full(), false)

	if actualValue, ok := queue.Dequeue(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	assert(queue.Size(), 0)

	assert(queue.Empty(), true)
	assert(queue.Full(), false)
	assert(len(queue.Values()), 0)
}

func TestQueueDequeueFull(t *testing.T) {
	assert := func(actualValue interface{}, expectedValue interface{}) {
		if actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}

	queue := circularbuffer.New[int](2)
	assert(queue.Empty(), true)
	assert(queue.Full(), false)
	assert(queue.Size(), 0)

	queue.Enqueue(1)
	assert(queue.Size(), 1)

	queue.Enqueue(2)
	assert(queue.Size(), 2)
	assert(queue.Full(), true)
	if actualValue, ok := queue.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	queue.Enqueue(3) // overwrites 1
	assert(queue.Size(), 2)

	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, expectedValue := queue.Size(), 1; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, ok := queue.Peek(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := queue.Size(), 1; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	assert(queue.Size(), 0)

	if actualValue, ok := queue.Dequeue(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	assert(queue.Empty(), true)
	assert(queue.Full(), false)
	assert(len(queue.Values()), 0)
}

func TestQueueIteratorOnEmpty(t *testing.T) {
	queue := circularbuffer.New[int](3)
	it := queue.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty queue")
	}
}

func TestQueueIteratorNext(t *testing.T) {
	queue := circularbuffer.New[string](3)
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")

	it := queue.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
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

	queue.Clear()
	it = queue.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty queue")
	}
}

func TestQueueIteratorPrev(t *testing.T) {
	queue := circularbuffer.New[string](3)
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")

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
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
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

func TestQueueIteratorBegin(t *testing.T) {
	queue := circularbuffer.New[string](3)
	it := queue.Iterator()
	it.Begin()
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestQueueIteratorEnd(t *testing.T) {
	queue := circularbuffer.New[string](3)
	it := queue.Iterator()

	if index := it.Index(); index != -1 {
		t.Errorf("Got %v expected %v", index, -1)
	}

	it.End()
	if index := it.Index(); index != 0 {
		t.Errorf("Got %v expected %v", index, 0)
	}

	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	it.End()
	if index := it.Index(); index != queue.Size() {
		t.Errorf("Got %v expected %v", index, queue.Size())
	}

	it.Prev()
	if index, value := it.Index(), it.Value(); index != queue.Size()-1 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, queue.Size()-1, "c")
	}
}

func TestQueueIteratorFirst(t *testing.T) {
	queue := circularbuffer.New[string](3)
	it := queue.Iterator()
	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestQueueIteratorLast(t *testing.T) {
	queue := circularbuffer.New[string](3)
	it := queue.Iterator()
	if actualValue, expectedValue := it.Last(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 2 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "c")
	}
}

func TestQueueIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// NextTo (empty)
	{
		queue := circularbuffer.New[string](3)
		it := queue.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty queue")
		}
	}

	// NextTo (not found)
	{
		queue := circularbuffer.New[string](3)
		queue.Enqueue("xx")
		queue.Enqueue("yy")
		it := queue.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty queue")
		}
	}

	// NextTo (found)
	{
		queue := circularbuffer.New[string](3)
		queue.Enqueue("aa")
		queue.Enqueue("bb")
		queue.Enqueue("cc")
		it := queue.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty queue")
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

func TestQueueIteratorPrevTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	// PrevTo (empty)
	{
		queue := circularbuffer.New[string](3)
		it := queue.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty queue")
		}
	}

	// PrevTo (not found)
	{
		queue := circularbuffer.New[string](3)
		queue.Enqueue("xx")
		queue.Enqueue("yy")
		it := queue.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty queue")
		}
	}

	// PrevTo (found)
	{
		queue := circularbuffer.New[string](3)
		queue.Enqueue("aa")
		queue.Enqueue("bb")
		queue.Enqueue("cc")
		it := queue.Iterator()
		it.End()
		if !it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty queue")
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

func TestQueueIterator(t *testing.T) {
	assert := func(actualValue interface{}, expectedValue interface{}) {
		if actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}

	queue := circularbuffer.New[string](2)

	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c") // overwrites "a"

	it := queue.Iterator()

	if actualIndex, expectedIndex := it.Index(), -1; actualIndex != expectedIndex {
		t.Errorf("Got %v expected %v", actualIndex, expectedIndex)
	}

	assert(it.Next(), true)

	if actualValue, actualIndex, expectedValue, expectedIndex := it.Value(), it.Index(), "b", 0; actualValue != expectedValue || actualIndex != expectedIndex {
		t.Errorf("Got %v expected %v, Got %v expected %v", actualValue, expectedValue, actualIndex, expectedIndex)
	}

	assert(it.Next(), true)

	if actualValue, actualIndex, expectedValue, expectedIndex := it.Value(), it.Index(), "c", 1; actualValue != expectedValue || actualIndex != expectedIndex {
		t.Errorf("Got %v expected %v, Got %v expected %v", actualValue, expectedValue, actualIndex, expectedIndex)
	}

	assert(it.Next(), false)

	if actualIndex, expectedIndex := it.Index(), 2; actualIndex != expectedIndex {
		t.Errorf("Got %v expected %v", actualIndex, expectedIndex)
	}

	assert(it.Next(), false)

	assert(it.Prev(), true)

	if actualValue, actualIndex, expectedValue, expectedIndex := it.Value(), it.Index(), "c", 1; actualValue != expectedValue || actualIndex != expectedIndex {
		t.Errorf("Got %v expected %v, Got %v expected %v", actualValue, expectedValue, actualIndex, expectedIndex)
	}

	assert(it.Prev(), true)

	if actualValue, actualIndex, expectedValue, expectedIndex := it.Value(), it.Index(), "b", 0; actualValue != expectedValue || actualIndex != expectedIndex {
		t.Errorf("Got %v expected %v, Got %v expected %v", actualValue, expectedValue, actualIndex, expectedIndex)
	}

	assert(it.Prev(), false)

	if actualIndex, expectedIndex := it.Index(), -1; actualIndex != expectedIndex {
		t.Errorf("Got %v expected %v", actualIndex, expectedIndex)
	}
}

func TestQueueSerialization(t *testing.T) {
	queue := circularbuffer.New[string](3)
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")

	var err error
	assert := func() {
		testutils.SameElements(t, queue.Values(), []string{"a", "b", "c"})
		if actualValue, expectedValue := queue.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
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

func TestQueueString(t *testing.T) {
	c := circularbuffer.New[int](3)
	c.Enqueue(1)
	if !strings.HasPrefix(c.String(), "CircularBuffer") {
		t.Errorf("String should start with container name")
	}
}

func benchmarkEnqueue(b *testing.B, queue *circularbuffer.Queue[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Enqueue(n)
		}
	}
}

func benchmarkDequeue(b *testing.B, queue *circularbuffer.Queue[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Dequeue()
		}
	}
}

func BenchmarkArrayQueueDequeue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := circularbuffer.New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := circularbuffer.New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := circularbuffer.New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := circularbuffer.New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := circularbuffer.New[int](3)
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := circularbuffer.New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := circularbuffer.New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := circularbuffer.New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}
