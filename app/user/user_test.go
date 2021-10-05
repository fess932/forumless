package user

import (
	"forumless/app/models"
	repomock "forumless/app/repo/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForum_CreateUser(t *testing.T) {
	repo := &repomock.Iface{}

	// repository mock setup
	{
		repo.On("CreateUser", models.User{Name: "SomeName"}).Return(nil).Once()           // ok
		repo.On("CreateUser", models.User{Name: "SomeName"}).Return(models.ErrUserExists) // ok
		repo.On("CreateUser", models.User{}).Return(models.ErrUserEmpty)
	}

	usrRepo := New(repo)

	type args struct {
		usr models.User
		err error
	}

	tests := []struct {
		name string
		user *User
		args args
	}{
		{
			"ok",
			usrRepo,
			args{models.User{Name: "SomeName"}, nil},
		},
		{
			"err empty username",
			usrRepo,
			args{models.User{}, models.ErrUserEmpty},
		},
		{
			"err user exist",
			usrRepo,
			args{models.User{Name: "SomeName"}, models.ErrUserExists},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.user.CreateUser(tt.args.usr)
			assert.ErrorIs(t, tt.args.err, err)

		})
	}
}
