package forum

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"forumless/app/models"
	repomock "forumless/app/repo/mock"
	"forumless/app/user"
)

func TestForum_CreatePost(t *testing.T) {
	repo := &repomock.Iface{}

	// repository mock setup
	{
		repo.On("CreatePost", models.User{ID: 1}, models.Post{Text: "kek"}).Return(nil) // ok
		repo.On("CreatePost", models.User{}, models.Post{Text: "kek"}).Return(models.ErrUserNotFound)
		repo.On("CreatePost", models.User{ID: 1}, models.Post{}).Return(models.ErrWrongText)
		repo.On("CreatePost", models.User{}, models.Post{}).Return(models.ErrWrongText)
	}

	forum := New("", "", "", repo, user.New(repo))

	type args struct {
		usr  models.User
		post models.Post
		err  error
	}

	tests := []struct {
		name string
		f    *Forum
		args args
	}{
		{
			"ok",
			forum,
			args{models.User{ID: 1}, models.Post{Text: "kek"}, nil},
		},
		{
			"wrong user id",
			forum,
			args{models.User{}, models.Post{Text: "kek"}, models.ErrUserNotFound},
		},
		{
			"wrong text",
			forum,
			args{models.User{}, models.Post{}, models.ErrWrongText},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.f.CreatePost(tt.args.usr, tt.args.post)
			assert.ErrorIs(t, tt.args.err, err)

		})
	}
}

func TestForum_CreateComment(t *testing.T) {
	repo := &repomock.Iface{}

	// repository mock setup
	{
		repo.On("CreateComment", models.User{ID: 1}, models.Post{ID: 1}, models.Comment{Text: "com kek"}).Return(nil) // ok
		repo.On("CreateComment", models.User{ID: 1}, models.Post{ID: 1}, models.Comment{}).Return(models.ErrWrongText)
		repo.On("CreateComment", models.User{ID: 1}, models.Post{}, models.Comment{Text: "com kek"}).Return(models.ErrPostNotFound)
		repo.On("CreateComment", models.User{}, models.Post{}, models.Comment{Text: "com kek"}).Return(models.ErrUserNotFound)
	}

	forum := New("", "", "", repo, user.New(repo))

	type args struct {
		usr models.User
		pst models.Post
		cmm models.Comment
		err error
	}

	tests := []struct {
		name string
		frm  *Forum
		args args
	}{
		{
			"ok",
			forum,
			args{models.User{ID: 1}, models.Post{ID: 1}, models.Comment{Text: "com kek"}, nil},
		},
		{
			"empty comment",
			forum,
			args{models.User{ID: 1}, models.Post{ID: 1}, models.Comment{}, models.ErrWrongText},
		},
		{
			"post emtpy or wrong id",
			forum,
			args{models.User{ID: 1}, models.Post{}, models.Comment{Text: "com kek"}, models.ErrPostNotFound},
		},
		{
			"wrong user id",
			forum,
			args{models.User{}, models.Post{}, models.Comment{Text: "com kek"}, models.ErrUserNotFound},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.frm.CreateComment(tt.args.usr, tt.args.pst, tt.args.cmm)
			assert.ErrorIs(t, tt.args.err, err)

		})
	}
}
