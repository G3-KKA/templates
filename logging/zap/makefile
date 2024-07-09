WORKSPACE ?= $(shell pwd)
include ${WORKSPACE}/.env
export $(shell sed 's/=.*//' .env)
export WORKSPACE
default:
	go build -o ${WORKSPACE}/bin/main ${WORKSPACE}/cmd/main.go
	${WORKSPACE}/bin/main