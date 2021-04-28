p=$(shell pwd)

all: server

server:
	PROJ_DIR=$p go run ./main.go server

migrate:
	PROJ_DIR=$p go run ./main.go migrate

go-vendor:
	go mod tidy
