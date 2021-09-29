package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type PostgresRepo struct {
	db *pgx.Conn
}

func New(connstr string) *PostgresRepo {
	db, err := pgx.Connect(context.Background(), connstr)

	if err != nil {
		log.Println("ok, just run docker next time")
		// log.Fatal(err)
	}

	return &PostgresRepo{db}
}

func (pg PostgresRepo) GetName() string {
	return "default name"
}
