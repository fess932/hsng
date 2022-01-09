package mong

import (
	"context"
	"github.com/fess932/hsng/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Repo struct {
	context   context.Context
	rwTimeout time.Duration
	db        *mongo.Database
}

func (r Repo) GetUsers() []*model.User {

	ctx, cancel := context.WithTimeout(r.context, r.rwTimeout)
	defer cancel()

	collection := r.db.Collection("users")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("find all in users", err)
	}

	defer cur.Close(ctx)

	var users []*model.User
	for cur.Next(ctx) {
		var usr model.User
		err := cur.Decode(&usr)
		if err != nil {
			log.Fatal("decode usr: ", err)
		}
		users = append(users, &usr)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func (r Repo) GetUser(s string) *model.User {
	panic("implement me")
}

func New(ctx context.Context, rwTimeout time.Duration, db *mongo.Database) *Repo {
	return &Repo{
		context:   ctx,
		rwTimeout: rwTimeout,
		db:        db,
	}
}
