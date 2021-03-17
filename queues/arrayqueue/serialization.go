package arrayqueue

import "github.com/emirpasic/gods/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Queue)(nil)
	var _ containers.JSONDeserializer = (*Queue)(nil)
}

// ToJSON outputs the JSON representation of the queue.
func (queue *Queue) ToJSON() ([]byte, error) {
	return queue.list.ToJSON()
}

// FromJSON populates the queue from the input JSON representation.
func (queue *Queue) FromJSON(data []byte) error {
	return queue.list.FromJSON(data)
}
