generate:
	go run github.com/99designs/gqlgen generate

run:
	go run .

init:
	go run github.com/99designs/gqlgen init

PHONY: init generate

