package forum

import (
	"testing"

	"github.com/stretchr/testify/require"

	"forumless/app/models"
	repomock "forumless/app/repo/mock"
	"forumless/app/user"
)

func TestForum_CreatePost(t *testing.T) {
	repo := &repomock.Iface{}
	// repo.On("CreatePost", , mock.IsType(models.Post{})).Return(nil)

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
			args{models.User{ID: 1}, models.Post{Author: 1, Text: "kek"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// repo.On()
			repo.On("CreatePost", tt.args.u, tt.args.p)

			require.NoError(t, tt.f.CreatePost(tt.args.u, tt.args.p))
		})
	}
}
