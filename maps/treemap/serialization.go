// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"encoding/json"

	"github.com/emirpasic/gods/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Map)(nil)
	var _ containers.JSONDeserializer = (*Map)(nil)
	var _ json.Marshaler = (*Map)(nil)
	var _ json.Unmarshaler = (*Map)(nil)
}

// ToJSON outputs the JSON representation of the map.
func (m *Map) ToJSON() ([]byte, error) {
	return m.tree.ToJSON()
}

// FromJSON populates the map from the input JSON representation.
func (m *Map) FromJSON(data []byte) error {
	return m.tree.FromJSON(data)
}

// @implements json.Unmarshaler
func (m *Map) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// @implements json.Marshaler
func (m *Map) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
