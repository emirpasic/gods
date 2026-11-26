// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"bytes"
	"encoding/json"

	"github.com/emirpasic/gods/v2/containers"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Map[string, int])(nil)
var _ containers.JSONDeserializer = (*Map[string, int])(nil)

// ToJSON outputs the JSON representation of map.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)

	buf.WriteRune('{')

	it := m.Iterator()
	lastIndex := m.Size() - 1
	index := 0

	for it.Next() {
		km, err := json.Marshal(it.Key())
		if err != nil {
			return nil, err
		}
		buf.Write(km)

		buf.WriteRune(':')

		vm, err := json.Marshal(it.Value())
		if err != nil {
			return nil, err
		}
		buf.Write(vm)

		if index != lastIndex {
			buf.WriteRune(',')
		}

		index++
	}

	buf.WriteRune('}')

	return buf.Bytes(), nil
}

// FromJSON populates map from the input JSON representation.
func (m *Map[K, V]) FromJSON(data []byte) error {
	elements := make(map[K]V)
	if err := json.Unmarshal(data, &elements); err != nil {
		return err
	}

	// Read only top-level keys in source order. Searching the input for a
	// marshaled key can match a value or a nested key, and misses equivalent
	// escaped spellings. Keep the values decoded above so duplicate keys retain
	// encoding/json's map-value semantics.
	decoder := json.NewDecoder(bytes.NewReader(data))
	if _, err := decoder.Token(); err != nil {
		return err
	}
	keys := make([]K, 0, len(elements))
	seen := make(map[K]bool, len(elements))
	var keyMap map[K]json.RawMessage
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		name := token.(string)
		var value json.RawMessage
		if err := decoder.Decode(&value); err != nil {
			return err
		}

		// Let encoding/json convert object names to K, including named integer
		// types and keys implementing encoding.TextUnmarshaler.
		encodedName, err := json.Marshal(name)
		if err != nil {
			return err
		}
		member := append([]byte{'{'}, encodedName...)
		member = append(member, ':', 'n', 'u', 'l', 'l', '}')
		clear(keyMap)
		if err := json.Unmarshal(member, &keyMap); err != nil {
			return err
		}
		for key := range keyMap {
			if !seen[key] {
				keys = append(keys, key)
				seen[key] = true
			}
		}
	}

	m.Clear()
	for _, key := range keys {
		m.Put(key, elements[key])
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
