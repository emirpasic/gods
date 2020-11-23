// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import "github.com/emirpasic/gods/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Heap)(nil)
	var _ containers.JSONDeserializer = (*Heap)(nil)
}

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}

// MarshalJSON Implements json.Marshaler inerface
func (heap *Heap) MarshalJSON() ([]byte, error) {
	return heap.ToJSON()
}

// UnmarshalJSON Implements json.Unmarshaler inerface
func (heap *Heap) UnmarshalJSON(data []byte) error {
	return heap.FromJSON(data)
}
