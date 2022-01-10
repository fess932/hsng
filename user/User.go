package user

import (
	"github.com/fess932/hsng/graph/model"
)

type Usecase struct {
	repo Repo
}

func (u Usecase) CreateUser(user *model.User) *model.User {
	return u.repo.SaveUser(user)
}

func (u Usecase) AddFriend() (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

type Repo interface {
	GetUser(string) *model.User
	SaveUser(*model.User) *model.User
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
