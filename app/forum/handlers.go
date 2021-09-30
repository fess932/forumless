package forum

import (
	"encoding/json"
	"fmt"
	"forumless/app/models"
	"log"
	"net/http"
)

func (f Forum) MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to forum %s", f.Name)
}

type postReq struct {
	Text string
}

func (f Forum) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var p postReq

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Fatalln(err)
	}

	f.CreatePost(models.User{}, models.Post{Text: p.Text})
}

func (f Forum) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	f.CreateUser(models.User{}, models.Post{})
}
