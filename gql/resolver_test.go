package gql

import (
	"context"
	"errors"
	"github.com/DrewBurkhart/directory/model"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type MockDB struct {
	collection *mongo.Collection
}

func (mockDB MockDB) GetPractitioners(string) ([]*model.Practitioner, error) {
	return []*model.Practitioner{{ID: "test-id"}}, errors.New("test-error")
}

func TestPractitioners(t *testing.T) {
	r := &queryResolver{
		Resolver: &Resolver{&MockDB{}},
	}

	practitioners, err := r.Practitioners(context.TODO(), "test")

	if practitioners[0].ID != "test-id" {
		t.Errorf("GetPractitioners() got = %v, want test-id", practitioners[0].ID)
	}
	if err.Error() != "test-error" {
		t.Errorf("GetPractitioners() got = %v, want test-error", err.Error())
	}
}
