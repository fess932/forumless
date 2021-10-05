package repo

import "forumless/app/models"

type Iface interface {
	CreatePost(user models.User, post models.Post) error
	CreateComment(user models.User, comment models.Comment)
	CreateUser(name string) error
}
