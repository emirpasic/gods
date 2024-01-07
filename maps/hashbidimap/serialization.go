// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashbidimap

import (
	"encoding/json"

	"github.com/emirpasic/gods/v2/containers"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Map[string, int])(nil)
var _ containers.JSONDeserializer = (*Map[string, int])(nil)

// ToJSON outputs the JSON representation of the map.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	return m.forwardMap.ToJSON()
}

// FromJSON populates the map from the input JSON representation.
func (m *Map[K, V]) FromJSON(data []byte) error {
	var elements map[K]V
	err := json.Unmarshal(data, &elements)
	if err != nil {
		return err
	}

	m.Clear()
	for k, v := range elements {
		m.Put(k, v)
	}

	return nil
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *Map[K, V]) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
