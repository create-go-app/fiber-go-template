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

2. Run database and apply migrations (by default, for PostgreSQL):

```bash
make migration-up user=<db_user> host=<db_host> table=<db_table>
```

3. Rename `.env.example` to `.env` and fill it with your environment values.

4. Run project by this command:

```bash
make run
```

## 📦 Used packages

| Name                                                    | Version   | Type       |
| ------------------------------------------------------- | --------- | ---------- |
| [gofiber/fiber](https://github.com/gofiber/fiber)       | `v2.5.0`  | core       |
| [gofiber/jwt](https://github.com/gofiber/jwt)           | `v2.1.0`  | middleware |
| [stretchr/testify](https://github.com/stretchr/testify) | `v1.7.0`  | tests      |
| [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) | `v3.2.0`  | auth       |
| [joho/godotenv](https://github.com/joho/godotenv)       | `v1.3.0`  | config     |
| [jmoiron/sqlx](https://github.com/jmoiron/sqlx)         | `v1.3.1`  | database   |
| [jackc/pgx](https://github.com/jackc/pgx)               | `v4.10.1` | database   |
| [google/uuid](https://github.com/google/uuid)           | `v1.2.0`  | utils      |

## 🗄 Template structure

### ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project
- `./app/validators` folder for describe validators for models fields

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
SERVER_HOST="0.0.0.0"
SERVER_PORT=5000
SERVER_EMAIL="no-reply@example.com"
SERVER_EMAIL_PASSWORD="secret"

# JWT settings:
JWT_SECRET_TOKEN="secret"

# Database type:
DATABASE_TYPE="postgres"

# PostgreSQL settings:
POSTGRES_SERVER_URL="host=localhost dbname=postgres sslmode=disable"
POSTGRES_MAX_CONNECTIONS=100
POSTGRES_MAX_IDLE_CONNECTIONS=10
POSTGRES_MAX_LIFETIME_CONNECTIONS=2

# SSL settings:
LETS_ENCRYPT_EMAIL="mail@gmail.com"
DOMAIN_WITHOUT_WWW="example.com"
DOMAIN_WITH_WWW="www.example.com"

# SMTP severs settings:
SMTP_SERVER="smtp.example.com"
SMTP_PORT=25
```

## ⚠️ License

MIT &copy; [Vic Shóstak](https://github.com/koddr) & [True web artisans](https://1wa.co/).
