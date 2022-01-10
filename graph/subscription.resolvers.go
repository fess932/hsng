package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/fess932/hsng/graph/generated"
	"github.com/fess932/hsng/graph/model"
)

func (r *subscriptionResolver) LastTodo(ctx context.Context) (<-chan *model.Todo, error) {
	ch := make(chan *model.Todo)

	go func(ch chan *model.Todo) {
		t := time.NewTicker(time.Second * 3)

		for v := range t.C {
			ch <- &model.Todo{
				ID:   "123",
				Text: "time: " + v.String(),
			}
		}

	}(ch)

	return ch, nil
}

func (r *subscriptionResolver) LastMessage(ctx context.Context) (<-chan *model.Message, error) {
	userID := strconv.Itoa(rand.Int())

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			log.Println("CONTEXT LAST MESSAGE DONE")
			r.messager.Unsubscribe(userID)
		}
	}(ctx)
	return r.messager.Subscribe(userID), nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
