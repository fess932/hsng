package mong

import (
	"context"
	"github.com/fess932/hsng/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var mong = "mongodb://localhost:27017"
var rwTimeout = time.Second * 10

type MRepo struct {
	client  *mongo.Client
	db      *mongo.Database
	context context.Context
}

func New(ctx context.Context) *MRepo {
	if eMongPath := os.Getenv("MONGO_PATH"); eMongPath != "" {
		mong = eMongPath
	}

	log.Println("mongo patH:", mong)

	// context for conecting up
	tctx, cancel := context.WithTimeout(ctx, rwTimeout)
	defer cancel()

	// connect to db
	client, err := mongo.Connect(tctx, options.Client().ApplyURI(mong))
	go func(ctx context.Context, client *mongo.Client) {
		<-ctx.Done()
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
		log.Println("mong disconected")
	}(ctx, client)

	// ping db
	tctx, cancel = context.WithTimeout(context.Background(), rwTimeout)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("ping: ", err)
		panic(err)
	}
	log.Println("mongodb connect successful")

	return &MRepo{
		client:  client,
		db:      client.Database("testing"),
		context: ctx,
	}
}

func (mr *MRepo) ReadAllMessages() []*model.Message {
	ctx, cancel := context.WithTimeout(mr.context, rwTimeout)
	defer cancel()

	collection := mr.db.Collection("messages")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("find all in messages", err)
	}

	defer cur.Close(ctx)

	var msgs []*model.Message
	for cur.Next(ctx) {
		var msg *model.Message
		err := cur.Decode(msg)
		if err != nil {
			log.Fatal("decode msg: ", err)
		}
		msgs = append(msgs, msg)

		//bson.D{}
		//
		//result.
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return msgs
}

func (mr *MRepo) SaveMessage(message *model.Message) {
	ctx, cancel := context.WithTimeout(mr.context, rwTimeout)
	defer cancel()
	collection := mr.db.Collection("messages")

	res, err := collection.InsertOne(ctx, bson.D{{"text", message.Text}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inserted id: ", res.InsertedID)
}
