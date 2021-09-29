package main

import (
	"log"
	"sync"

	"github.com/golang/mock/gomock"

	"forumless/app/config"
	"forumless/app/forum"
	"forumless/app/repo/mock"
)

func main() {
	NewServer(config.New()).Run()
}

func NewServer(conf config.Config) *server {

	// repo := postgres.New("conn string")

	repo := mock.NewMockRepo(&gomock.Controller{})

	forums := []*forum.Forum{}

	for _, v := range conf.Forums {
		forums = append(forums, forum.New(v.Host, v.Name, v.Port, repo))
	}

	return &server{forums: forums}
}

type server struct {
	forums []*forum.Forum
	sync.WaitGroup
}

func (s server) Run() {
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
