package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/fess932/hsng/graph/model"
	"github.com/fess932/hsng/message"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos    []*model.Todo
	messager Messager
}

type Messager interface {
	SendMessage(*model.Message)
	GetMessages() []*model.Message
	Subscribe(string) <-chan *model.Message
	Unsubscribe(string)
}

func NewResolver() *Resolver {
	return &Resolver{
		messager: message.New(),
	}
}
