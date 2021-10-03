package user

type Userer interface {
	CreateUser(string) error
}

type User struct {
	repo Userer
}

func New(repo Userer) *User {
	return &User{repo}
}

func (u *User) CreateUser(name string) error {
	return u.repo.CreateUser(name)
}
