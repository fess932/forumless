include .env
.DEFAULT_GOAL := start

PROJECTNAME=$(shell basename "$(PWD)")

start:
	go run ./app/main.go

migration-create:
	migrate create -ext sql -dir db/migrations -seq initial

migration-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migration-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

mock:
	mockgen.exe -source .\app\forum\forum.go -destination .\app\repo\mock\repo.go -package mock