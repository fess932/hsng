package user

import (
	"github.com/fess932/hsng/graph/model"
)

type Usecase struct {
	repo Repo
}

type Repo interface {
	GetUser(string) *model.User
	GetUsers() []*model.User
}

func New(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u Usecase) GetUser(id string) *model.User {
	panic("implement me")
}

func (u Usecase) GetUsers() []*model.User {
	return u.repo.GetUsers()
}
