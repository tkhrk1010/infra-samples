package models

import (
	"fmt"
	"errors"

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
func (id *UserAccountId) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"value": id.value,
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

// Equals compares other UserAccountId.
func (ua *UserAccountId) Equals(other *UserAccountId) bool {
	return ua.value == other.value
}

// snapshotから集約を復元するときに使う
func ConvertUserAccountIdFromJSON(value map[string]interface{}) mo.Result[UserAccountId] {
	return NewUserAccountIdFromString(value["value"].(string))
}
