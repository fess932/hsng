package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/fess932/hsng/graph/generated"
	"github.com/fess932/hsng/graph/model"
)

func (r *messageResolver) Sender(ctx context.Context, obj *model.Message) (*model.User, error) {
	return &model.User{
		ID:   "senderID",
		Name: "reciver name",
	}, nil
}

func (r *messageResolver) Reciver(ctx context.Context, obj *model.Message) (*model.User, error) {
	return &model.User{
		ID:   "reciverID",
		Name: "reciver name",
	}, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) SendMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	message := &model.Message{
		Text:      input.Text,
		SenderID:  "expapmle",
		ReciverID: "exapmle",
	}

	r.messager.SendMessage(message)

	return message, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return r.messager.GetMessages(), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.user.GetUsers(), nil
}

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

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type messageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
