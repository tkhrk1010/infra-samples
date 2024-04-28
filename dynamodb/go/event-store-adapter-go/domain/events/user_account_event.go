package events

import (
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
)

const (
	EventTypeUserAccountCreated     = "UserAccountCreated"
	EventTypeUserAccountNameChanged = "UserAccountNameChanged"
)

// UserAccountEvent is a domain event for user account.
type UserAccountEvent interface {
	esag.Event
	ToJSON() map[string]interface{}
}
