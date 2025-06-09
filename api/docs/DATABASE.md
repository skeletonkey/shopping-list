# Database Setup

This project uses SQLite as the database with [pressly/goose](https://github.com/pressly/goose) for database migrations.

## Database Schema

The application creates three main tables:

### Family Table

- `id` - Primary key (auto-increment)
- `uuid` - Unique identifier for external use
- `name` - Family name
- `display_name` - Optional display name

### List Table

- `id` - Primary key (auto-increment)
- `uuid` - Unique identifier for external use
- `name` - List name
- `display_name` - Optional display name
- `family_id` - Foreign key to family table

### Item Table

- `id` - Primary key (auto-increment)
- `uuid` - Unique identifier for external use
- `item` - Item name
- `list_id` - List ID (0 for global items, specific list ID for list-specific items)

## Setup Instructions

### Goose Environmental Variables

It is assumed that goose commands will be run from the `api` directory.

- GOOSE_DRIVER - The database driver to use
  - `export GOOSE_DRIVER=sqlite3`
- GOOSE_DBSTRING - The database connection string
  - `export GOOSE_DBSTRING=./database/dev.db`
- GOOSE_MIGRATION_DIR - The directory containing the migration files (default: .)
  - `export GOOSE_MIGRATION_DIR=./database/goose`

### Current Migrations

- `00001_initial_db.sql` - Creates the initial family, list, and item tables with indexes

### Adding New Migrations

1. Create a new SQL: `goose -s create add_some_column sql`
2. Add up/down SQL statements to file
3. Run `goose up`

## Database Utilities

### View Schema

```bash
sqlite3 database/dev.db ".schema"
```

### View Tables

```bash
sqlite3 database/dev.db ".tables"
```

### Query Data

```bash
sqlite3 database/dev.db "SELECT * FROM family;"
```

## Dependencies

- `github.com/pressly/goose/v3` - Database migration tool
- `github.com/mattn/go-sqlite3` - SQLite driver for Go

## Configuration

Database configuration is managed through the config file (`config/dev.json`):

```json
{
  "db": {
    "location": "database/dev.db"
  }
}
