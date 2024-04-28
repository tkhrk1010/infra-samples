package events

import (
	"fmt"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
)

type UserAccountNameChanged struct {
	Id          string
	AggregateId esag.AggregateId
	TypeName    string
	SeqNr       uint64
	ExecutorId  models.UserAccountId
	Name        string
	OccurredAt  uint64
}

func NewUserAccountNameChanged(id string, aggregateId esag.AggregateId, seqNr uint64, name string, occurredAt uint64) *UserAccountNameChanged {
	return &UserAccountNameChanged{
		Id:          id,
		AggregateId: aggregateId,
		TypeName:    "UserAccountNameChanged",
		SeqNr:       seqNr,
		Name:        name,
		OccurredAt:  occurredAt,
	}
}

func (e *UserAccountNameChanged) String() string {
	return fmt.Sprintf("UserAccountNameChanged{Id: %s, AggregateId: %s, SeqNr: %d, Name: %s, OccurredAt: %d}", e.Id, e.AggregateId, e.SeqNr, e.Name, e.OccurredAt)
}

func (e *UserAccountNameChanged) GetId() string {
	return e.Id
}

func (e *UserAccountNameChanged) GetTypeName() string {
	return e.TypeName
}

func (e *UserAccountNameChanged) GetAggregateId() esag.AggregateId {
	return e.AggregateId
}

func (e *UserAccountNameChanged) GetSeqNr() uint64 {
	return e.SeqNr
}

func (e *UserAccountNameChanged) GetOccurredAt() uint64 {
	return e.OccurredAt
}

func (e *UserAccountNameChanged) IsCreated() bool {
	return false
}

func (g *UserAccountNameChanged) GetExecutorId() *models.UserAccountId {
	return &g.ExecutorId
}

func (g *UserAccountNameChanged) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"type_name":    g.GetTypeName(),
		"id":           g.Id,
		"aggregate_id": g.AggregateId,
		"name":         g.Name,
		"executor_id":  g.ExecutorId.ToJSON(),
		"seq_nr":       g.SeqNr,
		"occurred_at":  g.OccurredAt,
	}
}