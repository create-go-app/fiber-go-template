# ./platform

**Folder with platform specific functionality**. This directory contains all the platform-level logic that will build up the actual project, like setting up the database, cache server instance or store for migrations.

- `./platform/database` folder with database configuration (by default, PostgreSQL)
- `./platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)
