package main

import (
	"fmt"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
)

const (
	EventTypeUserAccountCreated     = "UserAccountCreated"
	EventTypeUserAccountNameChanged = "UserAccountNameChanged"
)

// UserAccountEvent is a domain event for user account.
type UserAccountEvent interface {
	esag.Event
	GetExecutorId() *UserAccountId
	ToJSON() map[string]interface{}
}

type UserAccountCreated struct {
	Id          string
	AggregateId esag.AggregateId
	TypeName    string
	SeqNr       uint64
	ExecutorId  UserAccountId
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

func (g *UserAccountCreated) GetExecutorId() *UserAccountId {
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

// ---

type UserAccountNameChanged struct {
	Id          string
	AggregateId esag.AggregateId
	TypeName    string
	SeqNr       uint64
	ExecutorId  UserAccountId
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

func (g *UserAccountNameChanged) GetExecutorId() *UserAccountId {
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