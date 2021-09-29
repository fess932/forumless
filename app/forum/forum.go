package forum

import "log"

type Repo interface {
	GetName() string
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

func (f Forum) GetName() {
	f.repo.GetName()
}
