all: build run

install:
	@go mod vendor

test:
	@echo "Ini test"

build: test
	@echo "Hello make"
	@echo "do something"
	@go build -o output main.go

run:
	@./output
	