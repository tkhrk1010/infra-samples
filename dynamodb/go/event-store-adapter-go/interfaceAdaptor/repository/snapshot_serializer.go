package repository

import (
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain"
	"encoding/json"

	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
)

type SnapshotSerializer struct{}

func NewSnapshotSerializer() *SnapshotSerializer {
	return &SnapshotSerializer{}
}

func (s *SnapshotSerializer) Serialize(aggregate esag.Aggregate) ([]byte, error) {
	result, err := json.Marshal(aggregate.(*domain.UserAccount).ToJSON())
	if err != nil {
		return nil, esag.NewSerializationError("Failed to serialize the snapshot", err)
	}
	return result, nil
}

func (s *SnapshotSerializer) Deserialize(data []byte, aggregateMap *map[string]interface{}) error {
	err := json.Unmarshal(data, aggregateMap)
	if err != nil {
		return esag.NewDeserializationError("Failed to deserialize the snapshot", err)
	}
	return nil
}
