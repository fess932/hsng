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

func New(ctx context.Context, rwTimeout time.Duration, db *mongo.Database) *Repo {
	return &Repo{
		db:        db,
		rwTimeout: rwTimeout,
		context:   ctx,
	}
}

func (r *Repo) ReadAllMessages() []*model.Message {
	ctx, cancel := context.WithTimeout(r.context, r.rwTimeout)
	defer cancel()

	collection := r.db.Collection("messages")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("find all in messages", err)
	}

	defer cur.Close(ctx)

	var msgs []*model.Message
	for cur.Next(ctx) {
		var msg model.Message
		err := cur.Decode(&msg)
		if err != nil {
			log.Fatal("decode msg: ", err)
		}
		msgs = append(msgs, &msg)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return msgs
}

func (r *Repo) SaveMessage(message *model.Message) {
	ctx, cancel := context.WithTimeout(r.context, r.rwTimeout)
	defer cancel()
	collection := r.db.Collection("messages")

	res, err := collection.InsertOne(ctx, bson.D{{"text", message.Text}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inserted id: ", res.InsertedID)
}
