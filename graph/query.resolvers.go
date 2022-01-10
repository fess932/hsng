package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/fess932/hsng/graph/generated"
	"github.com/fess932/hsng/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return r.messager.GetMessages(), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.user.GetUsers(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
