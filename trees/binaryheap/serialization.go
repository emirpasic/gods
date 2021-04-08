// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import (
	"encoding/json"

	"github.com/emirpasic/gods/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Heap)(nil)
	var _ containers.JSONDeserializer = (*Heap)(nil)
	var _ json.Marshaler = (*Heap)(nil)
	var _ json.Unmarshaler = (*Heap)(nil)
}

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}

// @implements json.Unmarshaler
func (heap *Heap) UnmarshalJSON(bytes []byte) error {
	return heap.FromJSON(bytes)
}

// @implements json.Marshaler
func (heap *Heap) MarshalJSON() ([]byte, error) {
	return heap.ToJSON()
}
