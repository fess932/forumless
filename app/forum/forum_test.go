package forum

import (
	"testing"

	"forumless/app/models"
	repomock "forumless/app/repo/mock"
	"forumless/app/user"
)

func TestForum_CreatePost(t *testing.T) {
	repo := &repomock.Iface{}

	type args struct {
		u models.User
		p models.Post
	}

	tests := []struct {
		name string
		f    *Forum
		args args
	}{
		{
			"sample",
			New("", "", "", repo, user.New(repo)),
			args{models.User{}, models.Post{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.CreatePost(tt.args.u, tt.args.p)
		})
	}
}
