include .env

PROJECTNAME=$(shell basename "$(PWD)")

start:
	go run ./app/main.go

migration-create:
	migrate create -ext sql -dir db/migrations -seq todo_placeholder

migration-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up