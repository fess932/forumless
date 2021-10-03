package forum

import (
	"fmt"
	"forumless/app/models"
	"forumless/app/user"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Repo interface {
	CreatePost(models.User, models.Post) error
}

type Forum struct {
	Host string
	Name string
	Port string

	repo Repo

	user *user.User
}

func New(host, name, port string, repo Repo, user *user.User) *Forum {
	return &Forum{host, name, port, repo, user}
}

func (f Forum) Run() {
	r := chi.NewRouter()

	// forum
	{
		r.Get("/", f.MainHandler)
		r.Post("/post", f.CreatePostHandler)
		r.Post("/user", f.CreateUserHandler)
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

func (f Forum) CreateUser(u models.User, p models.Post) {
	if err := f.repo.CreatePost(u, p); err != nil {
		log.Fatal(err)
	}
}
