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

// CreatePostHandler godoc
// @Summary Create post
// @Description create post by user and text
// @ID create-post
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
//#@Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
//#@Failure 400,404 {object} httputil.HTTPError
//#@Failure 500 {object} httputil.HTTPError
//#@Failure default {object} httputil.DefaultError
// @Router /accounts/{id} [get]
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
