package forum

import (
	"testing"

	"forumless/app/models"
	"forumless/app/repo/mock"

	"github.com/golang/mock/gomock"
)

func TestForum_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)

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
			New("", "", "", mock.NewMockRepo(ctrl)),
			args{models.User{}, models.Post{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.CreatePost(tt.args.u, tt.args.p)
		})
	}
}
