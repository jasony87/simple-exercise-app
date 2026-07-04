# Makefile for simple-exercise-app
.PHONY: generate api run build clean
all: generate api run

@PHONY: generate
generate:
	go generate ./...

@PHONY: run
run:
	go run cmd/main.go

@PHONY: build
build:
	go build -o simple-exercise-app cmd/main.go

@PHONY: clean
clean:
	rm -f simple-exercise-app