package events

import (
	"fmt"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
)

type UserAccountCreated struct {
	// TODO: scopeをprivateにする
	Id          string
	AggregateId esag.AggregateId
	TypeName    string
	SeqNr       uint64
	ExecutorId  models.UserAccountId
	Name        string
	OccurredAt  uint64
}

func NewUserAccountCreated(id string, aggregateId esag.AggregateId, seqNr uint64, name string, occurredAt uint64) *UserAccountCreated {
	return &UserAccountCreated{
		Id:          id,
		AggregateId: aggregateId,
		TypeName:    "UserAccountCreated",
		SeqNr:       seqNr,
		Name:        name,
		OccurredAt:  occurredAt,
	}
}

func (g *UserAccountCreated) ToJSON() map[string]interface{} {
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


func (e *UserAccountCreated) String() string {
	return fmt.Sprintf("UserAccountCreated{Id: %s, AggregateId: %s, SeqNr: %d, Name: %s, OccurredAt: %d}", e.Id, e.AggregateId, e.SeqNr, e.Name, e.OccurredAt)
}

func (e *UserAccountCreated) GetId() string {
	return e.Id
}

func (e *UserAccountCreated) GetTypeName() string {
	return e.TypeName
}

func (e *UserAccountCreated) GetAggregateId() esag.AggregateId {
	return e.AggregateId
}

func (g *UserAccountCreated) GetExecutorId() *models.UserAccountId {
	return &g.ExecutorId
}

func (e *UserAccountCreated) GetSeqNr() uint64 {
	return e.SeqNr
}

func (e *UserAccountCreated) GetOccurredAt() uint64 {
	return e.OccurredAt
}

func (e *UserAccountCreated) IsCreated() bool {
	return true
}
