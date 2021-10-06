package user

import (
	"encoding/json"
	"forumless/app/models"
	"net/http"
)

// CreateUserHandler godoc
// @Summary Create user
// @Description create post by user and text
// @ID create-user
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
//#@Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
//#@Failure 400,404 {object} httputil.HTTPError
//#@Failure 500 {object} httputil.HTTPError
//#@Failure default {object} httputil.DefaultError
// @Router /user [post]
func (u *User) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Name string `json:"name"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := u.CreateUser(models.User{Name: req.Name}); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
