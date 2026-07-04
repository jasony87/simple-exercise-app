# Makefile for simple-exercise-app
.PHONY: all generate run build clean

all: generate run

generate:
	go generate ./...

run:
	go run cmd/main.go

build:
	go build -o simple-exercise-app cmd/main.go

clean:
	rm -f simple-exercise-app
