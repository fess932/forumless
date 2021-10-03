package postgres

import (
	"context"
	"forumless/app/models"
	"log"

	"github.com/jackc/pgx/v4"
)

type Repo struct {
	db *pgx.Conn
}

func (pg Repo) CreatePost(user models.User, post models.Post) error {
	_, err := pg.db.Exec(context.Background(),
		"INSERT INTO forumless.public.post (data, user_id) VALUES($1, $2)", post.Text, user.ID,
	)

	return err
}

func (pg Repo) CreateComment(user models.User, comment models.Comment) error {
	panic("implement me")
}

func (pg Repo) CreateUser(name string) error {
	_, err := pg.db.Exec(context.Background(),
		`INSERT INTO forumless.public."user" (name) VALUES($1)`, name,
	)

	return err
}

func New(connstr string) *Repo {
	db, err := pgx.Connect(context.Background(), connstr)

	if err != nil {
		log.Println("ok, just run docker next time")
		log.Fatal(err)
	}

	return &Repo{db}
}

func (pg Repo) GetName() string {
	return "default name"
}
