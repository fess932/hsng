include .env
export

generate:
	go run github.com/99designs/gqlgen generate
	go run .

run:
	go run .

init:
	go run github.com/99designs/gqlgen init

PHONY: init generate run
