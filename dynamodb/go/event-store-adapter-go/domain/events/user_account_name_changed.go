package events

import (
	"fmt"
	"time"
	"github.com/oklog/ulid/v2"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
)

type UserAccountNameChanged struct {
	id          string
	aggregateId models.UserAccountId
	name        string
	seqNr       uint64
	executorId  models.UserAccountId
	occurredAt  uint64
}

func NewUserAccountNameChanged(aggregateId models.UserAccountId, name string, seqNr uint64, executorId models.UserAccountId) UserAccountNameChanged {
	id := ulid.Make().String()
	now := time.Now()
	occurredAt := uint64(now.UnixNano() / 1e6)
	return UserAccountNameChanged{id, aggregateId, name, seqNr, executorId, occurredAt}
}

func (e *UserAccountNameChanged) String() string {
	return fmt.Sprintf("UserAccountNameChanged{Id: %s, AggregateId: %s, SeqNr: %d, Name: %s, OccurredAt: %d}", e.id, e.aggregateId, e.seqNr, e.name, e.occurredAt)
}

func (e *UserAccountNameChanged) GetId() string {
	return e.id
}

func (e *UserAccountNameChanged) GetTypeName() string {
	return "UserAccountNameChanged"
}

func (e *UserAccountNameChanged) GetName() *string {
	return &e.name
}

func (e *UserAccountNameChanged) GetAggregateId() esag.AggregateId {
	return &e.aggregateId
}

func (e *UserAccountNameChanged) GetSeqNr() uint64 {
	return e.seqNr
}

func (e *UserAccountNameChanged) GetOccurredAt() uint64 {
	return e.occurredAt
}

func (e *UserAccountNameChanged) IsCreated() bool {
	return false
}

func (e *UserAccountNameChanged) GetExecutorId() *models.UserAccountId {
	return &e.executorId
}

// NewUserAccountNameChangedFrom is a constructor for UserAccountNameChanged
func NewUserAccountNameChangedFrom(id string, aggregateId models.UserAccountId, name string, seqNr uint64, executorId models.UserAccountId, occurredAt uint64) UserAccountNameChanged {
	return UserAccountNameChanged{id, aggregateId, name, seqNr, executorId, occurredAt}
}

func (e *UserAccountNameChanged) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"type_name":    e.GetTypeName(),
		"id":           e.id,
		"aggregate_id": e.aggregateId,
		"name":         e.name,
		"executor_id":  e.executorId.ToJSON(),
		"seq_nr":       e.seqNr,
		"occurred_at":  e.occurredAt,
	}
}