.PHONY: clean test security build run

BUILD_DIR = $(PWD)/build
APP_NAME = apiserver

clean:
	rm -rf ./build

test:
	go test -cover ./...

security:
	gosec -quiet ./...

build: security test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: clean build
	$(BUILD_DIR)/$(APP_NAME)
