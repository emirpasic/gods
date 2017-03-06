// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import "github.com/emirpasic/gods/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Map)(nil)
	var _ containers.JSONDeserializer = (*Map)(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (m *Map) ToJSON() ([]byte, error) {
	return m.tree.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (m *Map) FromJSON(data []byte) error {
	return m.tree.FromJSON(data)
}
