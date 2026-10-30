// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import (
	"github.com/emirpasic/gods/v2/containers"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Heap[int])(nil)
var _ containers.JSONDeserializer = (*Heap[int])(nil)

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap[T]) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
// Decoded values are ordered using the heap's comparator.
func (heap *Heap[T]) FromJSON(data []byte) error {
	if err := heap.list.FromJSON(data); err != nil {
		return err
	}
	for i := heap.list.Size()/2 - 1; i >= 0; i-- {
		heap.bubbleDownIndex(i)
	}
	return nil
}

// UnmarshalJSON @implements json.Unmarshaler
func (heap *Heap[T]) UnmarshalJSON(bytes []byte) error {
	return heap.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (heap *Heap[T]) MarshalJSON() ([]byte, error) {
	return heap.ToJSON()
}
