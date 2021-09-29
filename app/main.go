package main

import (
	"forumless/app/forum"
	"log"
)

func main() {
	log.Print("lel")

	NewServer().Run()
}

func NewServer() *server {
	return &server{}
}

type server struct {
	forum forum.Forum
}

func (s server) Run() {

}
