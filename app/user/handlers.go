package user

import (
	"encoding/json"
	"forumless/app/models"
	"net/http"
)

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
