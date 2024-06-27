WORKDIR ?= $(shell pwd)/..
include ./.env
export $(shell sed 's/=.*//' .env)

default:
	go build -o ${WORKDIR}/bin/main ${WORKDIR}/cmd/main.go
	${WORKDIR}/bin/main