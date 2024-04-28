package repository

import (
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/events"
	"encoding/json"

	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
)

type EventSerializer struct{}

func NewEventSerializer() *EventSerializer {
	return &EventSerializer{}
}

func (s *EventSerializer) Serialize(event esag.Event) ([]byte, error) {
	result, err := json.Marshal(event.(events.UserAccountEvent).ToJSON())
	if err != nil {
		return nil, esag.NewSerializationError("Failed to serialize the event", err)
	}
	return result, nil
}

func (s *EventSerializer) Deserialize(data []byte, eventMap *map[string]interface{}) error {
	err := json.Unmarshal(data, eventMap)
	if err != nil {
		return esag.NewDeserializationError("Failed to deserialize the event", err)
	}
	return nil
}
