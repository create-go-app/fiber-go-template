include .env
.PHONY: clean test security build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://${DB_USER}:${DB_PASS}@${DB_HOST}/${DB_NAME}?sslmode=${DB_SSL}
SERVER_PORT = `echo $(SERVER_URL) | cut -d ":" -f2` # extract port from server url defined in .env

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

test: security
	go test -cover ./...

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag.init build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

migrate.create:
	migrate create -ext sql -dir $(MIGRATIONS_FOLDER) $(name)

migrate.version:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" version

migrate.goto:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" goto $(version)

docker.build:
	docker build -t fiber-go-template .

docker.run: docker.fiber docker.postgres

docker.stop:
	docker stop dev-fiber dev-postgres

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber: docker.network
	docker run --rm -d \
		--name dev-fiber \
		--network dev-network \
		-p $(SERVER_PORT):$(SERVER_PORT) \
		fiber-go-template

docker.postgres: docker.network
	docker run --rm -d \
		--name dev-postgres \
		--network dev-network \
		-e POSTGRES_USER=${DB_USER} \
		-e POSTGRES_PASSWORD=${DB_PASS} \
		-e POSTGRES_DB=${DB_NAME} \
		-e POSTGRES_PORT=${DB_PORT} \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p ${DB_PORT}:${DB_PORT} \
		postgres

swag.init:
	swag init