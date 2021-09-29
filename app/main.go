package main

import (
	"forumless/app/config"
	"forumless/app/forum"
	"sync"
)

func main() {
	NewServer(config.New()).Run()
}

func NewServer(conf config.Config) *server {
	forums := []forum.Forum{}

	for _, v := range conf.Forums {
		forums = append(forums, forum.Forum{Host: v.Host, Name: v.Name, Port: v.Port})
	}

	return &server{forums: forums}
}

type server struct {
	forums []forum.Forum
	sync.WaitGroup
}

func (s server) Run() {
	s.Add(len(s.forums))

	for _, f := range s.forums {
		f := f // hack for save var in closure
		go func(f forum.Forum) {
			f.Run()
			s.Done()
		}(f)
	}

	s.Wait()
}
