package forum

import (
	"encoding/json"
	"fmt"
	"forumless/app/models"
	"log"
	"net/http"
	"strconv"
)

func (f Forum) MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to forum %s", f.Name)
}

type postReq struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (f Forum) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var p postReq

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Fatalln(err)
	}

	log.Println(p)

	s, err := strconv.Atoi(p.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if err := f.CreatePost(models.User{ID: s}, models.Post{Text: p.Text}); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (f Forum) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	f.CreateUser(models.User{}, models.Post{})
}
