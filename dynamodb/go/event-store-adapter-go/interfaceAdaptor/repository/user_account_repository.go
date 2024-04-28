package repository

import (
	"fmt"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
	// "encoding/json"

	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain"
	// "github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
	// "github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/events"
	
)

type UserAccountRepository struct {
	eventStore esag.EventStore
}


// func EventConverter(m map[string]interface{}) (esag.Event, error) {
// 	// TODO: impl
// 	return nil, nil
// }

// func SnapshotConverter(m map[string]interface{}) (esag.Aggregate, error) {
// 	userAccountId, err := models.ConvertUserAccountIdFromJSON(m["id"].(map[string]interface{})).Get()
// 	if err != nil {
// 		return nil, err
// 	}
// 	name := m["name"].(string)
// 	seqNr := uint64(m["seq_nr"].(float64))
// 	version := uint64(m["version"].(float64))
// 	result := domain.NewUserAccountFrom(userAccountId, name, seqNr, version)
// 	return &result, nil
// }

// type EventSerializer struct{}

// func NewEventSerializer() *EventSerializer {
// 	return &EventSerializer{}
// }

// func (s *EventSerializer) Serialize(event esag.Event) ([]byte, error) {
// 	result, err := json.Marshal(event.(events.UserAccountEvent).ToJSON())
// 	if err != nil {
// 		return nil, esag.NewSerializationError("Failed to serialize the event", err)
// 	}
// 	return result, nil
// }

// func (s *EventSerializer) Deserialize(data []byte, eventMap *map[string]interface{}) error {
// 	err := json.Unmarshal(data, eventMap)
// 	if err != nil {
// 		return esag.NewDeserializationError("Failed to deserialize the event", err)
// 	}
// 	return nil
// }


// type SnapshotSerializer struct{}

// func NewSnapshotSerializer() *SnapshotSerializer {
// 	return &SnapshotSerializer{}
// }

// func (s *SnapshotSerializer) Serialize(aggregate esag.Aggregate) ([]byte, error) {
// 	result, err := json.Marshal(aggregate.(*domain.UserAccount).ToJSON())
// 	if err != nil {
// 		return nil, esag.NewSerializationError("Failed to serialize the snapshot", err)
// 	}
// 	return result, nil
// }

// func (s *SnapshotSerializer) Deserialize(data []byte, aggregateMap *map[string]interface{}) error {
// 	err := json.Unmarshal(data, aggregateMap)
// 	if err != nil {
// 		return esag.NewDeserializationError("Failed to deserialize the snapshot", err)
// 	}
// 	return nil
// }



func NewUserAccountRepository(eventStore esag.EventStore) *UserAccountRepository {
	return &UserAccountRepository{
		eventStore: eventStore,
	}
}

func (r *UserAccountRepository) StoreEvent(event esag.Event, version uint64) error {
	return r.eventStore.PersistEvent(event, version)
}

func (r *UserAccountRepository) StoreEventAndSnapshot(event esag.Event, aggregate esag.Aggregate) error {
	return r.eventStore.PersistEventAndSnapshot(event, aggregate)
}

func (r *UserAccountRepository) FindById(id esag.AggregateId) (*domain.UserAccount, error) {
	result, err := r.eventStore.GetLatestSnapshotById(id)
	if err != nil {
		return nil, err
	}
	if result.Empty() {
		return nil, fmt.Errorf("not found")
	} else {
		events, err := r.eventStore.GetEventsByIdSinceSeqNr(id, result.Aggregate().GetSeqNr()+1)
		if err != nil {
			return nil, err
		}
		return domain.ReplayUserAccount(events, result.Aggregate().(*domain.UserAccount)), nil
	}
}


