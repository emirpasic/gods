package serialization

import (
	"encoding/json"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
)

// ListSerializationExample demonstrates how to serialize and deserialize lists to and from JSON
func ListSerializationExample() {
	list := arraylist.New()
	list.Add("a", "b", "c")

	// Serialization (marshalling)
	jsonBytes, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes)) // ["a","b","c"]

	// Deserialization (unmarshalling)
	jsonBytes = []byte(`["a","b"]`)
	err = json.Unmarshal(jsonBytes, list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list) // ArrayList ["a","b"]
}

// MapSerializationExample demonstrates how to serialize and deserialize maps to and from JSON
func MapSerializationExample() {
	m := hashmap.New()
	m.Put("a", "1")
	m.Put("b", "2")
	m.Put("c", "3")

	// Serialization (marshalling)
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes)) // {"a":"1","b":"2","c":"3"}

	// Deserialization (unmarshalling)
	jsonBytes = []byte(`{"a":"1","b":"2"}`)
	err = json.Unmarshal(jsonBytes, m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m) // HashMap {"a":"1","b":"2"}
}
