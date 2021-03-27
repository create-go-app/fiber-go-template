# ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like _setting up the database_ or _cache server instance_ and _storing migrations_.

- `./platform/cache` folder with in-memory cache setup functions
- `./platform/database` folder with database configuration
- `./platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)
