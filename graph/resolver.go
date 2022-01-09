package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"github.com/fess932/hsng/graph/model"
	"github.com/fess932/hsng/message"
	messageMongo "github.com/fess932/hsng/message/repo/mongo"
	"github.com/fess932/hsng/user"
	userMongo "github.com/fess932/hsng/user/repo/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos    []*model.Todo
	messager Messager
	user     User
}

type Messager interface {
	SendMessage(*model.Message)
	GetMessages() []*model.Message
	Subscribe(string) <-chan *model.Message
	Unsubscribe(string)
}

type User interface {
	GetUser(id string) *model.User
	GetUsers() []*model.User
}

func NewResolver(ctx context.Context) *Resolver {
	// default mongo store
	ms := NewMongoStore(ctx)

	return &Resolver{
		messager: message.New(ms.messageRepo),
		user:     user.New(ms.userRepo),
	}
}

/////////////////  MONGO STORE ENGINE //////////////////////////////
var mong = "mongodb://localhost:27017"
var rwTimeout = time.Second * 10 // default db ops timeout

type MongoStore struct {
	db     *mongo.Database
	client *mongo.Client

	messageRepo *messageMongo.Repo
	userRepo    *userMongo.Repo
}

func NewMongoStore(ctx context.Context) *MongoStore {
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
	db := client.Database("testing")
	return &MongoStore{
		db:     db,
		client: client,

		messageRepo: messageMongo.New(ctx, rwTimeout, db),
		userRepo:    userMongo.New(ctx, rwTimeout, db),
	}
}
