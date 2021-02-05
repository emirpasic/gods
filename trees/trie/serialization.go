package trie

import "github.com/emirpasic/gods/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Trie)(nil)
	var _ containers.JSONDeserializer = (*Trie)(nil)
}

// ToJSON outputs the JSON representation of the heap.
func (t *Trie) ToJSON() ([]byte, error) {
	elements := make([]interface{})
	it := trie.Iterator(t)
	for it.Next() {
		elements.append(it.Value())
	}
	return json.Marshal(&elements)
}

// FromJSON populates the tree from the input JSON representation.
func (t *Trie) FromJSON(data []byte) error {
	elements := make([]interface{})
	err := json.Unmarshal(data, &elements)
	if err == nil {
		for _, value := range elements {
			t.insert(value)
		}
	}
	return err
}
