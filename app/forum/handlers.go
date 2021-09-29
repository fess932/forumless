package forum

import (
	"forumless/app/models"
	"net/http"
)

func (f Forum) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	f.CreatePost(models.User{}, models.Post{})
}

func (f Forum) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	f.CreateUser(models.User{}, models.Post{})
}
