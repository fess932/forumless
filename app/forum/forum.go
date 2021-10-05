// этот модуль является точкой входа для остальных модулей, подключает rest обработчики и является местом
// регистрации новых пользователей

package forum

import (
	"fmt"
	"forumless/app/models"
	"forumless/app/repo"
	"forumless/app/user"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Forum struct {
	Host string
	Name string
	Port string

	repo repo.Forumer
	user *user.User
}

func New(host, name, port string, repo repo.Forumer, user *user.User) *Forum {
	return &Forum{host, name, port, repo, user}
}

func (f Forum) Run() {
	r := chi.NewRouter()

	// forum
	{

		r.Get("/", f.MainHandler)
		r.Post("/post", f.CreatePostHandler)
	}

	// user
	{
		r.Post("/user", f.user.CreateUserHandler)
	}

	host := fmt.Sprintf(":%s", f.Port)

	log.Printf("forum %s started at http://0.0.0.0:%s", f.Name, f.Port)
	log.Fatal(http.ListenAndServe(host, r))
}

func (f Forum) CreatePost(u models.User, p models.Post) error {
	return f.repo.CreatePost(u, p)
}

func (f Forum) CreateComment(u models.User, p models.Post, c models.Comment) error {
	return f.repo.CreateComment(u, p, c)
}
