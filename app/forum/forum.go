package forum

import (
	"forumless/app/models"
	"log"
)

type Repo interface {
	CreatePost(models.User, models.Post) error
}

type Forum struct {
	Host string
	Name string
	Port string

	repo Repo
}

func New(host, name, port string, repo Repo) *Forum {
	return &Forum{host, name, port, repo}
}

func (f Forum) Run() {
	log.Println("run fourm", f.Name)
}

func (f Forum) CreatePost(u models.User, p models.Post) {
	if err := f.repo.CreatePost(u, p); err != nil {
		log.Fatal(err)
	}
}

func (f Forum) CreateUser(u models.User, p models.Post) {
	if err := f.repo.CreatePost(u, p); err != nil {
		log.Fatal(err)
	}
}
