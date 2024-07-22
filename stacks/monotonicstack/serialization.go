package monotonicstack

import (
	"encoding/json"

	"github.com/emirpasic/gods/v2/containers"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Stack[int])(nil)
var _ containers.JSONDeserializer = (*Stack[int])(nil)

// ToJSON outputs the JSON representation of the stack.
func (stack *Stack[T]) ToJSON() ([]byte, error) {
	return json.Marshal(stack.Stack.Values())
}

// FromJSON populates the stack from the input JSON representation.
func (stack *Stack[T]) FromJSON(data []byte) error {
	var elements []T
	if err := json.Unmarshal(data, &elements); err != nil {
		return err
	}
	stack.Stack.Clear()
	for _, e := range elements {
		stack.Stack.Push(e)
	}
	return nil
}

// UnmarshalJSON @implements json.Unmarshaler
func (stack *Stack[T]) UnmarshalJSON(bytes []byte) error {
	return stack.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (stack *Stack[T]) MarshalJSON() ([]byte, error) {
	return stack.ToJSON()
}
