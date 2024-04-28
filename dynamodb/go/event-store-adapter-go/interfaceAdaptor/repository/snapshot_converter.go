package repository

import (
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"
)

func SnapshotConverter(m map[string]interface{}) (esag.Aggregate, error) {
	userAccountId, err := models.ConvertUserAccountIdFromJSON(m["id"].(map[string]interface{})).Get()
	if err != nil {
		return nil, err
	}
	name := m["name"].(string)
	seqNr := uint64(m["seq_nr"].(float64))
	version := uint64(m["version"].(float64))
	result := domain.NewUserAccountFrom(userAccountId, name, seqNr, version)
	return &result, nil
}
