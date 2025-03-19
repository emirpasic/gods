package deque

import (
	"testing"
)

func TestDeque(t *testing.T) {
	deque := NewDeque[int]()

	if !deque.IsEmpty() {
		t.Errorf("Expected deque to be empty")
	}

	deque.AddFront(1)
	deque.AddBack(2)
	deque.AddFront(0)

	if deque.Size() != 3 {
		t.Errorf("Expected deque size to be 3, got %d", deque.Size())
	}

	front, _ := deque.RemoveFront()
	if front != 0 {
		t.Errorf("Expected front item to be 0, got %v", front)
	}

	back, _ := deque.RemoveBack()
	if back != 2 {
		t.Errorf("Expected back item to be 2, got %v", back)
	}

	if deque.Size() != 1 {
		t.Errorf("Expected deque size to be 1, got %d", deque.Size())
	}
}

func TestDequeIterator(t *testing.T) {
	deque := NewDeque[int]()
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	it := deque.Iterator()
	expected := []int{1, 2, 3}
	for i := 0; it.Next(); i++ {
		item := it.Value()
		if item != expected[i] {
			t.Errorf("Expected item %v, got %v", expected[i], item)
		}
	}
}

func TestDequeSerialization(t *testing.T) {
	deque := NewDeque[int]()
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	data, err := deque.ToJSON()
	if err != nil {
		t.Errorf("Failed to serialize deque: %v", err)
	}

	newDeque := NewDeque[int]()
	err = newDeque.FromJSON(data)
	if err != nil {
		t.Errorf("Failed to deserialize deque: %v", err)
	}

	if newDeque.Size() != deque.Size() {
		t.Errorf("Expected deque size to be %d, got %d", deque.Size(), newDeque.Size())
	}

	it := newDeque.Iterator()
	expected := []int{1, 2, 3}
	for i := 0; it.Next(); i++ {
		item := it.Value()
		if item != expected[i] {
			t.Errorf("Expected item %v, got %v", expected[i], item)
		}
	}
}
