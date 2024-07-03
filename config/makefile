WORKSPACE ?= $(shell pwd)
include ${WORKSPACE}/.env
export $(shell sed 's/=.*//' .env)
export WORKSPACE
example:
	go build -o ${WORKSPACE}/bin/main ${WORKSPACE}/cmd/main.go
	${WORKSPACE}/bin/main