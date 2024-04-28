// 集約ルート
package domain

import (
	"fmt"
	"math/rand"
	"time"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
	"github.com/oklog/ulid/v2"

	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/events"
)

type UserAccount struct {
	id      models.UserAccountId
	name    string
	seqNr   uint64
	version uint64
}

// ToJSON converts the aggregate to JSON.
//
// However, this method is out of layer.
func (ua *UserAccount) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":       ua.id.ToJSON(),
		"name":     ua.name,
		"seq_nr":   ua.seqNr,
		"version":  ua.version,
	}
}


func NewUserAccount(id models.UserAccountId, name string) (*UserAccount, *events.UserAccountCreated) {
	aggregate := UserAccount{
		id:      id,
		name:    name,
		seqNr:   0,
		version: 1,
	}
	aggregate.seqNr += 1
	eventId := newULID()
	return &aggregate, events.NewUserAccountCreated(eventId.String(), &id, aggregate.seqNr, name, uint64(time.Now().UnixNano()))
}

func ReplayUserAccount(events []esag.Event, snapshot *UserAccount) *UserAccount {
	result := snapshot
	for _, event := range events {
		result = result.applyEvent(event)
	}
	return result
}

func (ua *UserAccount) applyEvent(event esag.Event) *UserAccount {
	switch e := event.(type) {
	case *events.UserAccountNameChanged:
		update, err := ua.Rename(e.Name)
		if err != nil {
			panic(err)
		}
		return update.Aggregate
	}
	return ua
}

func (ua *UserAccount) String() string {
	return fmt.Sprintf("UserAccount{Id: %s, Name: %s}", ua.id.String(), ua.name)
}

func (ua *UserAccount) GetId() esag.AggregateId {
	return &ua.id
}

func (ua *UserAccount) GetSeqNr() uint64 {
	return ua.seqNr
}

func (ua *UserAccount) GetVersion() uint64 {
	return ua.version
}

func (ua *UserAccount) WithVersion(version uint64) esag.Aggregate {
	result := *ua
	result.version = version
	return &result
}

type UserAccountResult struct {
	Aggregate *UserAccount
	Event     *events.UserAccountNameChanged
}

func (ua *UserAccount) Rename(name string) (*UserAccountResult, error) {
	updatedUserAccount := *ua
	updatedUserAccount.name = name
	updatedUserAccount.seqNr += 1
	event := events.NewUserAccountNameChanged(newULID().String(), &ua.id, updatedUserAccount.seqNr, name, uint64(time.Now().UnixNano()))
	return &UserAccountResult{&updatedUserAccount, event}, nil
}

func (ua *UserAccount) Equals(other *UserAccount) bool {
	return ua.id.Equals(&other.id) && ua.name == other.name && ua.seqNr == other.seqNr && ua.version == other.version 
}

// UUIDよりsort性能がいいULIDを使う
func newULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

func (ua *UserAccount) ChangeName(name string) (*UserAccount, *events.UserAccountNameChanged) {
	updatedUserAccount := *ua
	updatedUserAccount.name = name
	updatedUserAccount.seqNr += 1
	event := events.NewUserAccountNameChanged(newULID().String(), &ua.id, updatedUserAccount.seqNr, name, uint64(time.Now().UnixNano()))
	return &updatedUserAccount, event
}


// snapshotから集約を復元するときに使う
func NewUserAccountFrom(id models.UserAccountId, name string, seqNr uint64, version uint64) UserAccount {
	return UserAccount{id, name, seqNr, version}
}

