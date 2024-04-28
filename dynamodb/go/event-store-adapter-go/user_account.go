package main

import (
	"fmt"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
	"math/rand"
	"time"
	"errors"

	"github.com/oklog/ulid/v2"
	"github.com/samber/mo"
)

const UserAccountIdPrefix = "UserAccount"


type UserAccountId struct {
	value string
}

func NewUserAccountId(value string) UserAccountId {
	return UserAccountId{value: value}
}

func (id *UserAccountId) GetTypeName() string {
	return "UserAccountId"
}

func (id *UserAccountId) GetValue() string {
	return id.value
}

func (id *UserAccountId) String() string {
	return fmt.Sprintf("UserAccount{TypeName: %s, Valuie: %s}", id.GetTypeName(), id.value)
}

func (id *UserAccountId) AsString() string {
	return fmt.Sprintf("%s-%s", id.GetTypeName(), id.value)
}

// ToJSON converts to JSON.
//
// However, this method is out of layer.
func (g *UserAccountId) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"value": g.value,
	}
}

// NewUserAccountIdFromString is a constructor for UserAccountId
// It creates UserAccountId from string
func NewUserAccountIdFromString(value string) mo.Result[UserAccountId] {
	if value == "" {
		return mo.Err[UserAccountId](errors.New("UserAccountId is empty"))
	}
	if len(value) > len(UserAccountIdPrefix) && value[0:len(UserAccountIdPrefix)] == UserAccountIdPrefix {
		value = value[len(UserAccountIdPrefix)+1:]
	}
	return mo.Ok(UserAccountId{value: value})
}

// snapshotから集約を復元するときに使う
func ConvertUserAccountIdFromJSON(value map[string]interface{}) mo.Result[UserAccountId] {
	return NewUserAccountIdFromString(value["value"].(string))
}






type UserAccount struct {
	id      UserAccountId
	name    string
	seqNr   uint64
	version uint64
}

// ToJSON converts the aggregate to JSON.
//
// However, this method is out of layer.
func (g *UserAccount) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":       g.id.ToJSON(),
		"name":     g.name,
		"seq_nr":   g.seqNr,
		"version":  g.version,
	}
}


func NewUserAccount(id UserAccountId, name string) (*UserAccount, *UserAccountCreated) {
	aggregate := UserAccount{
		id:      id,
		name:    name,
		seqNr:   0,
		version: 1,
	}
	aggregate.seqNr += 1
	eventId := newULID()
	return &aggregate, NewUserAccountCreated(eventId.String(), &id, aggregate.seqNr, name, uint64(time.Now().UnixNano()))
}

func replayUserAccount(events []esag.Event, snapshot *UserAccount) *UserAccount {
	result := snapshot
	for _, event := range events {
		result = result.applyEvent(event)
	}
	return result
}

func (ua *UserAccount) applyEvent(event esag.Event) *UserAccount {
	switch e := event.(type) {
	case *UserAccountNameChanged:
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
	Event     *UserAccountNameChanged
}

func (ua *UserAccount) Rename(name string) (*UserAccountResult, error) {
	updatedUserAccount := *ua
	updatedUserAccount.name = name
	updatedUserAccount.seqNr += 1
	event := NewUserAccountNameChanged(newULID().String(), &ua.id, updatedUserAccount.seqNr, name, uint64(time.Now().UnixNano()))
	return &UserAccountResult{&updatedUserAccount, event}, nil
}

func (ua *UserAccount) Equals(other *UserAccount) bool {
	return ua.id.value == other.id.value && ua.name == other.name
}

func newULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

func (ua *UserAccount) ChangeName(name string) (*UserAccount, *UserAccountNameChanged) {
	updatedUserAccount := *ua
	updatedUserAccount.name = name
	updatedUserAccount.seqNr += 1
	event := NewUserAccountNameChanged(newULID().String(), &ua.id, updatedUserAccount.seqNr, name, uint64(time.Now().UnixNano()))
	return &updatedUserAccount, event
}


// snapshotから集約を復元するときに使う
func NewUserAccountFrom(id UserAccountId, name string, seqNr uint64, version uint64) UserAccount {
	return UserAccount{id, name, seqNr, version}
}

