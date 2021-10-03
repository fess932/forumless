package user

import "net/http"

func (u *User) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := u.CreateUser("KEKS"); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
