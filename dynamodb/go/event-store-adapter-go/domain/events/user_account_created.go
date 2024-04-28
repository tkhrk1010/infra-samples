package events

import (
	"fmt"
	"time"

	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
	"github.com/oklog/ulid/v2"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
)

type UserAccountCreated struct {
	id          string
	aggregateId models.UserAccountId
	name        string
	seqNr       uint64
	// TODO: UserAccountが別集約前提でコピペしてきたから、EmailIdとかが変な感じになってる。sampleのモデル考え直す必要あり。
	// TODO: EmailIdにモデル差し替え
	emailId    models.EmailId
	occurredAt uint64
}

func NewUserAccountCreated(aggregateId models.UserAccountId, name string, seqNr uint64, emailId models.EmailId) UserAccountCreated {
	id := ulid.Make().String()
	now := time.Now()
	occurredAt := uint64(now.UnixNano() / 1e6)
	return UserAccountCreated{id, aggregateId, name, seqNr, emailId, occurredAt}
}

// NewUserAccountCreatedFrom is a constructor for UserAccountCreated
func NewUserAccountCreatedFrom(id string, aggregateId models.UserAccountId, name string, seqNr uint64, emailId models.EmailId, occurredAt uint64) UserAccountCreated {
	return UserAccountCreated{id, aggregateId, name, seqNr, emailId, occurredAt}
}

func (u *UserAccountCreated) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"type_name":    u.GetTypeName(),
		"id":           u.id,
		"aggregate_id": u.aggregateId.ToJSON(),
		"name":         u.name,
		"email_id":     u.emailId.ToJSON(),
		"seq_nr":       u.seqNr,
		"occurred_at":  u.occurredAt,
	}
}

func (e *UserAccountCreated) String() string {
	return fmt.Sprintf("UserAccountCreated{Id: %s, aggregateId: %s, SeqNr: %d, Name: %s, OccurredAt: %d}", e.id, e.aggregateId, e.seqNr, e.name, e.occurredAt)
}

func (e *UserAccountCreated) GetId() string {
	return e.id
}

func (e *UserAccountCreated) GetTypeName() string {
	// dataとして持たずに固定値を返す振る舞いだけ持つ
	return "UserAccountCreated"
}

func (e *UserAccountCreated) GetAggregateId() esag.AggregateId {
	return &e.aggregateId
}

func (e *UserAccountCreated) GetEmailId() *models.EmailId {
	return &e.emailId
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
