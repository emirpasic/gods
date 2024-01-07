// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treeset

import (
	"encoding/json"

	"github.com/emirpasic/gods/v2/containers"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Set[int])(nil)
var _ containers.JSONDeserializer = (*Set[int])(nil)

// ToJSON outputs the JSON representation of the set.
func (set *Set[T]) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Set[T]) FromJSON(data []byte) error {
	var elements []T
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (set *Set[T]) UnmarshalJSON(bytes []byte) error {
	return set.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (set *Set[T]) MarshalJSON() ([]byte, error) {
	return set.ToJSON()
}
