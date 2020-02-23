# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m

# Define app variables
NAME=apiserver
BUILD=./build

.PHONY: all

all: clean build run

.PHONY: clean

clean:
	@rm -rf $(BUILD)
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend was cleaned!"

.PHONY: test

test:
	@go test ./...
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend was tested!"

.PHONY: check

check:
	@gosec ./...
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend was checked by gosec!"

.PHONY: run

run:
	@go run ./cmd/$(NAME)/...

.PHONY: build

build: clean check
	@CGO_ENABLED=0 GOARCH=amd64
	@GOOS=darwin go build -o $(BUILD)/darwin/$(NAME) ./cmd/$(NAME)/...
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for macOS x64 was builded!"
	@GOOS=linux go build -o $(BUILD)/linux/$(NAME) ./cmd/$(NAME)/...
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for GNU/Linux x64 was builded!"
	@GOOS=windows go build -ldflags="-H windowsgui" -o $(BUILD)/windows/$(NAME).exe ./cmd/$(NAME)/...
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for MS Windows x64 was builded!"
