include .env
.DEFAULT_GOAL := start

PROJECTNAME=$(shell basename "$(PWD)")

start:
	go run ./app/main.go

test:
	go test ./...

migration-create:
	migrate create -ext sql -dir db/migrations -seq initial

migration-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migration-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

mock:
	mockery --name=Iface --recursive --output=app/repo/mock

swagger:
	swag init -d app -g forum/forum.go