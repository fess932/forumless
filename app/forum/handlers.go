package forum

import (
	"fmt"
	"forumless/app/models"
	"net/http"
)

func (f Forum) MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to forum %s", f.Name)
}

func (f Forum) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	f.CreatePost(models.User{}, models.Post{})
}

func (f Forum) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	f.CreateUser(models.User{}, models.Post{})
}
