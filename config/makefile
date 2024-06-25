include ./.env

export $(shell sed 's/=.*//' .env)

default:
	go build -o ../bin/main ../cmd/main.go
	../bin/main