// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
.DEFAULT_GOAL=build

include Makefile.db


build:
	go fmt ./...
	go vet ./...
	go build -o bin/api app/*.go

install:
	cp bin/api /usr/local/sbin/api

golib-latest:
	go get -u github.com/pressly/goose/v3@latest
	go get -u github.com/labstack/echo/v4@latest
	go get -u github.com/mattn/go-sqlite3@latest
	go get -u github.com/skeletonkey/lib-core-go@latest
	go get -u github.com/skeletonkey/lib-instance-gen-go@latest

	go mod tidy

app-init:
	go generate
