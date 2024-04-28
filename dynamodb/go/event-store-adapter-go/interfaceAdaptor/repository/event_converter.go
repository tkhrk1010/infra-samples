package repository

import (
	"fmt"
	"log/slog"

	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/events"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"

	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
)

func EventConverter(m map[string]interface{}) (esag.Event, error) {
	slog.Info(fmt.Sprintf("EventConverter: %v", m))
	eventId := m["id"].(string)
	userAccountId, err := models.ConvertUserAccountIdFromJSON(m["aggregate_id"].(map[string]interface{})).Get()
	if err != nil {
		return nil, err
	}
	emailId, err := models.ConvertEmailIdFromJSON(m["email_id"].(map[string]interface{})).Get()
	if err != nil {
		return nil, err
	}
	seqNr := uint64(m["seq_nr"].(float64))
	occurredAt := uint64(m["occurred_at"].(float64))
	switch m["type_name"].(string) {
	case "UserAccountCreated":
		userAccountName := m["name"].(string)
		event := events.NewUserAccountCreatedFrom(
			eventId,
			userAccountId,
			userAccountName,
			seqNr,
			emailId,
			occurredAt,
		)
		return &event, nil
	case "UserAccountNameChanged":
		userAccountName := m["name"].(string)
		if err != nil {
			return nil, err
		}
		event := events.NewUserAccountNameChangedFrom(
			eventId,
			userAccountId,
			userAccountName,
			seqNr,
			emailId,
			occurredAt,
		)
		return &event, nil
	default:
		return nil, fmt.Errorf("unknown event type: %s", m["type_name"].(string))
	}
}
