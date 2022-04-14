// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package circularbuffer

import (
	"encoding/json"
	"github.com/emirpasic/gods/containers"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Queue)(nil)
var _ containers.JSONDeserializer = (*Queue)(nil)

// ToJSON outputs the JSON representation of queue's elements.
func (queue *Queue) ToJSON() ([]byte, error) {
	return json.Marshal(queue.values[:queue.maxSize])
}

// FromJSON populates list's elements from the input JSON representation.
func (queue *Queue) FromJSON(data []byte) error {
	var values []interface{}
	err := json.Unmarshal(data, &values)
	if err == nil {
		for _, value := range values {
			queue.Enqueue(value)
		}
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (queue *Queue) UnmarshalJSON(bytes []byte) error {
	return queue.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (queue *Queue) MarshalJSON() ([]byte, error) {
	return queue.ToJSON()
}
