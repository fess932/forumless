package main

import (
	"forumless/app/user"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"forumless/app/config"
	"forumless/app/forum"
	"forumless/app/repo/postgres"
)

func main() {
	NewServer(config.New()).Run()
}

func NewServer(conf config.Config) *server {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	repo := postgres.New(os.Getenv("POSTGRESQL_URL"))
	up := user.New(repo)

	var forums []*forum.Forum
	for _, v := range conf.Forums {
		forums = append(forums, forum.New(v.Host, v.Name, v.Port, repo, up))
	}

	return &server{forums: forums}
}

type server struct {
	forums []*forum.Forum
	sync.WaitGroup
}

func (s *server) Run() {
	s.Add(len(s.forums))

	for _, f := range s.forums {
		f := f // hack for save var in closure
		go func(f *forum.Forum) {
			f.Run()
			log.Println("forum done:", f.Name)
			s.Done()
		}(f)
	}

	s.Wait()
}
