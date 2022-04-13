// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedlistqueue

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	queue := New()
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if actualValue := queue.Values(); actualValue[0].(int) != 1 || actualValue[1].(int) != 2 || actualValue[2].(int) != 3 {
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
	queue := New()
	if actualValue, ok := queue.Peek(); actualValue != nil || ok {
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
	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()
	if actualValue, ok := queue.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := queue.Values(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}
}

func TestQueueIteratorOnEmpty(t *testing.T) {
	queue := New()
	it := queue.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty queue")
	}
}

func TestQueueIteratorNext(t *testing.T) {
	queue := New()
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

func TestQueueIteratorBegin(t *testing.T) {
	queue := New()
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

func TestQueueIteratorFirst(t *testing.T) {
	queue := New()
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

func TestQueueIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value interface{}) bool {
		return strings.HasSuffix(value.(string), "b")
	}

	// NextTo (empty)
	{
		queue := New()
		it := queue.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty queue")
		}
	}

	// NextTo (not found)
	{
		queue := New()
		queue.Enqueue("xx")
		queue.Enqueue("yy")
		it := queue.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty queue")
		}
	}

	// NextTo (found)
	{
		queue := New()
		queue.Enqueue("aa")
		queue.Enqueue("bb")
		queue.Enqueue("cc")
		it := queue.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty queue")
		}
		if index, value := it.Index(), it.Value(); index != 1 || value.(string) != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Next() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Index(), it.Value(); index != 2 || value.(string) != "cc" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "cc")
		}
		if it.Next() {
			t.Errorf("Should not go past last element")
		}
	}
}

func TestQueueSerialization(t *testing.T) {
	queue := New()
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")

	var err error
	assert := func() {
		if actualValue, expectedValue := fmt.Sprintf("%s%s%s", queue.Values()...), "abc"; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
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

	err = json.Unmarshal([]byte(`[1,2,3]`), &queue)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
}

func TestQueueString(t *testing.T) {
	c := New()
	c.Enqueue(1)
	if !strings.HasPrefix(c.String(), "LinkedListQueue") {
		t.Errorf("String should start with container name")
	}
}

func benchmarkEnqueue(b *testing.B, queue *Queue, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Enqueue(n)
		}
	}
}

func benchmarkDequeue(b *testing.B, queue *Queue, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Dequeue()
		}
	}
}

func BenchmarkArrayQueueDequeue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := New()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := New()
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}
