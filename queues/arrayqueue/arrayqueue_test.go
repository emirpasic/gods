package arrayqueue

import (
	"fmt"
	"testing"
)

func TestQueuePush(t *testing.T) {
	queue := New()
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

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
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestQueuePeek(t *testing.T) {
	queue := New()
	if actualValue, ok := queue.Peek(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if actualValue, ok := queue.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestQueuePop(t *testing.T) {
	queue := New()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Pop()
	if actualValue, ok := queue.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Pop(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Pop(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := queue.Pop(); actualValue != nil || ok {
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
	queue.Push("a")
	queue.Push("b")
	queue.Push("c")

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
}

func TestQueueIteratorPrev(t *testing.T) {
	queue := New()
	queue.Push("a")
	queue.Push("b")
	queue.Push("c")

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
	queue := New()
	it := queue.Iterator()
	it.Begin()
	queue.Push("a")
	queue.Push("b")
	queue.Push("c")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestQueueIteratorEnd(t *testing.T) {
	queue := New()
	it := queue.Iterator()

	if index := it.Index(); index != -1 {
		t.Errorf("Got %v expected %v", index, -1)
	}

	it.End()
	if index := it.Index(); index != 0 {
		t.Errorf("Got %v expected %v", index, 0)
	}

	queue.Push("a")
	queue.Push("b")
	queue.Push("c")
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
	queue := New()
	it := queue.Iterator()
	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	queue.Push("a")
	queue.Push("b")
	queue.Push("c")
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestQueueIteratorLast(t *testing.T) {
	queue := New()
	it := queue.Iterator()
	if actualValue, expectedValue := it.Last(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	queue.Push("a")
	queue.Push("b")
	queue.Push("c")
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 2 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "c")
	}
}

func TestQueueSerialization(t *testing.T) {
	queue := New()
	queue.Push("a")
	queue.Push("b")
	queue.Push("c")

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

	json, err := queue.ToJSON()
	assert()

	err = queue.FromJSON(json)
	assert()
}

func benchmarkPush(b *testing.B, queue *Queue, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Push(n)
		}
	}
}

func benchmarkPop(b *testing.B, queue *Queue, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Pop()
		}
	}
}

func BenchmarkArrayQueuePop100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := New()
	for n := 0; n < size; n++ {
		queue.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, queue, size)
}

func BenchmarkArrayQueuePop1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, queue, size)
}

func BenchmarkArrayQueuePop10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, queue, size)
}

func BenchmarkArrayQueuePop100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, queue, size)
}

func BenchmarkArrayQueuePush100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := New()
	b.StartTimer()
	benchmarkPush(b, queue, size)
}

func BenchmarkArrayQueuePush1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, queue, size)
}

func BenchmarkArrayQueuePush10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, queue, size)
}

func BenchmarkArrayQueuePush100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := New()
	for n := 0; n < size; n++ {
		queue.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, queue, size)
}
