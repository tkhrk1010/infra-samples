package events

import (
	"fmt"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
)

type UserAccountCreated struct {
	id          string
	aggregateId esag.AggregateId
	typeName    string
	seqNr       uint64
	executorId  models.UserAccountId
	name        string
	occurredAt  uint64
}

func NewUserAccountCreated(id string, aggregateId esag.AggregateId, seqNr uint64, name string, occurredAt uint64) *UserAccountCreated {
	return &UserAccountCreated{
		id:          id,
		aggregateId: aggregateId,
		typeName:    "UserAccountCreated",
		seqNr:       seqNr,
		name:        name,
		occurredAt:  occurredAt,
	}
}

func (g *UserAccountCreated) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"type_name":    g.GetTypeName(),
		"id":           g.id,
		"aggregate_id": g.aggregateId,
		"name":         g.name,
		"executor_id":  g.executorId.ToJSON(),
		"seq_nr":       g.seqNr,
		"occurred_at":  g.occurredAt,
	}
}


func (e *UserAccountCreated) String() string {
	return fmt.Sprintf("UserAccountCreated{Id: %s, aggregateId: %s, SeqNr: %d, Name: %s, OccurredAt: %d}", e.id, e.aggregateId, e.seqNr, e.name, e.occurredAt)
}

func (e *UserAccountCreated) GetId() string {
	return e.id
}

func (e *UserAccountCreated) GetTypeName() string {
	return e.typeName
}

func (e *UserAccountCreated) GetAggregateId() esag.AggregateId {
	return e.aggregateId
}

func (e *UserAccountCreated) GetExecutorId() *models.UserAccountId {
	return &e.executorId
}

func (e *UserAccountCreated) GetSeqNr() uint64 {
	return e.seqNr
}

func (e *UserAccountCreated) GetOccurredAt() uint64 {
	return e.occurredAt
}

func (e *UserAccountCreated) IsCreated() bool {
	return true
}
