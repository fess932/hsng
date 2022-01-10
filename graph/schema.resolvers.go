package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

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

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type messageResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
