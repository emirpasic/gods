// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package btree

import (
	"encoding/json"

	"github.com/emirpasic/gods/v2/containers"
	"github.com/emirpasic/gods/v2/internal/treeenc"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Tree[string, int])(nil)
var _ containers.JSONDeserializer = (*Tree[string, int])(nil)

// ToJSON outputs the JSON representation of the tree.
func (tree *Tree[K, V]) ToJSON() ([]byte, error) {
	elements := make(map[*treeenc.KeyMarshaler[K]]V)
	it := tree.Iterator()
	for it.Next() {
		elements[&treeenc.KeyMarshaler[K]{Key: it.Key()}] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the tree from the input JSON representation.
func (tree *Tree[K, V]) FromJSON(data []byte) error {
	elements := make(map[treeenc.KeyUnmarshaler[K]]V)
	err := json.Unmarshal(data, &elements)
	if err != nil {
		return err
	}

	tree.Clear()
	for ku, value := range elements {
		tree.Put(*ku.Key, value)
	}

	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (tree *Tree[K, V]) UnmarshalJSON(bytes []byte) error {
	return tree.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (tree *Tree[K, V]) MarshalJSON() ([]byte, error) {
	return tree.ToJSON()
}
