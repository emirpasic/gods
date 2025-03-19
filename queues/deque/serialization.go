package deque

import (
	"encoding/json"

	"github.com/emirpasic/gods/v2/containers"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Deque[int])(nil)
var _ containers.JSONDeserializer = (*Deque[int])(nil)

// ToJSON outputs the JSON representation of the deque.
func (deque *Deque[T]) ToJSON() ([]byte, error) {
	return json.Marshal(deque.items)
}

// FromJSON populates the deque from the input JSON representation.
func (deque *Deque[T]) FromJSON(data []byte) error {
	return json.Unmarshal(data, &deque.items)
}

// UnmarshalJSON @implements json.Unmarshaler
func (deque *Deque[T]) UnmarshalJSON(bytes []byte) error {
	return deque.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (deque *Deque[T]) MarshalJSON() ([]byte, error) {
	return deque.ToJSON()
}
