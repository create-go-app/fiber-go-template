# Fiber backend template for [Create Go App CLI](https://github.com/create-go-app/cli)

<img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;<a href="https://gocover.io/github.com/create-go-app/fiber-go-template/pkg/apiserver" target="_blank"><img src="https://img.shields.io/badge/Go_Cover-87%25-success?style=for-the-badge&logo=none" alt="go cover" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/fiber-go-template" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-mit-red?style=for-the-badge&logo=none" alt="license" />

[Fiber](https://gofiber.io/) is an Express.js inspired web framework build on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for **fast** development with **zero memory allocation** and **performance** in mind.

## ⚡️ Quick start

1. Create a new project with Fiber:

```bash
cgapp create

# Choose a backend framework:
#   net/http
# > Fiber
```

2. Run Docker container with database (_by default, for PostgreSQL_):

```bash
make docker.postgres
```

3. Apply migrations:

```bash
make migration.up user=<db_user> pass=<db_pass> host=<db_host> table=<db_table>
```

4. Rename `.env.example` to `.env` and fill it with your environment values.

5. Run project by this command:

```bash
make run
```

6. Go to API Docs page (Swagger): [localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html).

## 📦 Used packages

| Name                                                                  | Version   | Type       |
| --------------------------------------------------------------------- | --------- | ---------- |
| [gofiber/fiber](https://github.com/gofiber/fiber)                     | `v2.5.0`  | core       |
| [gofiber/jwt](https://github.com/gofiber/jwt)                         | `v2.1.0`  | middleware |
| [arsmn/fiber-swagger](https://github.com/arsmn/fiber-swagger)         | `v2.3.0`  | middleware |
| [stretchr/testify](https://github.com/stretchr/testify)               | `v1.7.0`  | tests      |
| [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)               | `v3.2.0`  | auth       |
| [joho/godotenv](https://github.com/joho/godotenv)                     | `v1.3.0`  | config     |
| [jmoiron/sqlx](https://github.com/jmoiron/sqlx)                       | `v1.3.1`  | database   |
| [jackc/pgx](https://github.com/jackc/pgx)                             | `v4.10.1` | database   |
| [swaggo/swag](https://github.com/swaggo/swag)                         | `v1.7.0`  | utils      |
| [google/uuid](https://github.com/google/uuid)                         | `v1.2.0`  | utils      |
| [go-playground/validator](https://github.com/go-playground/validator) | `v10.4.1` | utils      |

## 🗄 Template structure

### ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project
- `./app/validators` folder for describe validators for models fields

### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.

### ./pkg

**Folder with project-specific functionality**. This directory contains all the project-specific code tailored only for your business use case, like _configs_, _middleware_, _routes_ or _utils_.

- `./pkg/configs` folder for configuration functions
- `./pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `./pkg/routes` folder for describe routes of your project
- `./pkg/utils` folder with utility functions (server starter, error checker, etc)

### ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like _setting up the database_ or _cache server instance_ and _storing migrations_.

- `./platform/database` folder with database setup functions (by default, PostgreSQL)
- `./platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)

## ⚙️ Configuration

```ini
# .env

# Server settings:
SERVER_URL="0.0.0.0:5000"
SERVER_EMAIL="no-reply@example.com"
SERVER_EMAIL_PASSWORD="secret"

# JWT settings:
JWT_SECRET_KEY="secret"
JWT_REFRESH_KEY="refresh"

# Database settings:
DB_SERVER_URL="host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable"
DB_MAX_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10
DB_MAX_LIFETIME_CONNECTIONS=2

# SMTP severs settings:
SMTP_SERVER="smtp.example.com"
SMTP_PORT=25
```

## ⚠️ License

MIT &copy; [Vic Shóstak](https://github.com/koddr) & [True web artisans](https://1wa.co/).
