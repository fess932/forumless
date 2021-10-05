package user

import (
	"forumless/app/models"
	"forumless/app/repo"
)

type User struct {
	repo repo.Userer
}

func New(repo repo.Userer) *User {
	return &User{repo}
}

func (u *User) CreateUser(user models.User) error {
	return u.repo.CreateUser(user)
}
