package models

import (
	"errors"
	"fmt"
	"github.com/samber/mo"
)

const EmailIdPrefix = "Email"

// EmailId is a value object that represents a email id.
type EmailId struct {
	value string
}

// NewEmailId is the constructor for EmailId with generating id.
func NewEmailId(id string) EmailId {
	return EmailId{value: id}
}

// NewEmailIdFromString is the constructor for EmailId.
func NewEmailIdFromString(value string) mo.Result[EmailId] {
	if value == "" {
		return mo.Err[EmailId](errors.New("EmailId is empty"))
	}
	if len(value) > len(EmailIdPrefix) && value[0:len(EmailIdPrefix)] == EmailIdPrefix {
		value = value[len(EmailIdPrefix)+1:]
	}
	return mo.Ok(EmailId{value: value})
}

// ConvertEmailIdFromJSON is a constructor for EmailId.
func ConvertEmailIdFromJSON(value map[string]interface{}) mo.Result[EmailId] {
	return NewEmailIdFromString(value["value"].(string))
}

// ToJSON converts to JSON.
//
// However, this method is out of layer.
func (u *EmailId) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"value": u.value,
	}
}

func (u *EmailId) GetValue() string {
	return u.value
}

func (u *EmailId) GetTypeName() string {
	return "UserAccount"
}

func (u *EmailId) AsString() string {
	return fmt.Sprintf("%s-%s", u.GetTypeName(), u.GetValue())
}

func (u *EmailId) String() string {
	return u.AsString()
}

func (u *EmailId) Equals(other *EmailId) bool {
	return u.value == other.value
}
