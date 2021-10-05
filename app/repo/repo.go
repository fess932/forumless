package repo

import "forumless/app/models"

type Iface interface {
	Forumer
	Userer
}

type Forumer interface {
	CreatePost(user models.User, post models.Post) error
	CreateComment(user models.User, comment models.Comment) error
}

type Userer interface {
	CreateUser(user models.User) error
}
