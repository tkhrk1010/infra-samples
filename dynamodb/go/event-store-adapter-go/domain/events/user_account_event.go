package events

import (
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
)

const (
	EventTypeUserAccountCreated     = "UserAccountCreated"
	EventTypeUserAccountNameChanged = "UserAccountNameChanged"
)

// UserAccountEvent is a domain event for user account.
type UserAccountEvent interface {
	esag.Event
	GetExecutorId() *models.UserAccountId
	ToJSON() map[string]interface{}
}
