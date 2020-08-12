package db

import (
	"context"
	"github.com/DrewBurkhart/directory/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DB interface {
	GetPractitioners(domain string) ([]*model.Practitioner, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func New(client *mongo.Client) *MongoDB {
	practitioners := client.Database("practitioners").Collection("practitioners")
	return &MongoDB{
		collection: practitioners,
	}
}

func (db MongoDB) GetPractitioners(domain string) ([]*model.Practitioner, error) {
	res, err := db.collection.Find(context.TODO(), db.filter(domain))
	if err != nil {
		log.Printf("Error while fetching practitioners: %s", err.Error())
		return nil, err
	}
	var p []*model.Practitioner
	err = res.All(context.TODO(), &p)
	if err != nil {
		log.Printf("Error while decoding practitioners: %s", err.Error())
		return nil, err
	}
	return p, nil
}

func (db MongoDB) filter(domain string) bson.D {
	return bson.D{{
		"domains.name",
		bson.D{{
			"$regex",
			"^" + domain + ".*$",
		}, {
			"$options",
			"i",
		}},
	}}
}
