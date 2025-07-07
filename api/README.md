# Shopping List API

A simple, self-hosted shopping list API designed with Grug's philosophy of simplicity. This API powers a family-friendly shopping list application that works like paper but syncs between family members.

## Overview

This shopping list API is built for families who want a simple way to manage shared shopping lists without the complexity of modern apps. It follows the principle that "if you need to explain how to use it, it's already too complex."

### Key Features

- **Simple**: One screen does everything - add items, remove items, see the list
- **Family-Friendly**: Multiple family members can share lists via simple URLs
- **Self-Hosted**: Uses SQLite database for easy deployment on Raspberry Pi or any server
- **Multiple Lists**: Families can create and manage multiple lists (grocery, hardware, Christmas, etc.)
- **No Accounts**: No login required - security through URL obscurity
- **Works Like Paper**: Familiar interface that anyone can use without training

## Architecture

- **Backend**: Go with Echo v4 framework
- **Database**: SQLite with Goose migrations
- **Eventing**: Custom logging system for list item changes
- **Frontend**: Simple HTML/CSS/JavaScript (no frameworks)

## Quick Start

1. **Build and Run**:
   ```bash
   make build
   ./bin/shopping-list-api
   ```

2. **Access Your Family Lists**:
   - Navigate to `http://localhost:8080/api/v1/your-family-name`
   - Create lists and start adding items
   - Share the URL with family members

## URL Structure

The API follows a simple, intuitive URL pattern:

- `http://base_url/api/v1/family_name` - Family page showing all lists
- `http://base_url/api/v1/family_name/list_name` - Individual list page
- `http://base_url/api/v1/family_name/list_name/items` - API endpoint for list items

## Configuration

Families are configured via JSON configuration files. See the `config/` directory for examples.

Example family configuration:

```json
{
  "family": {
    "name": "smith",
    "display_name": "Smith Family"
  }
}
```

## Database

This application uses SQLite for simplicity and portability. Database migrations are handled by [pressly/goose](https://github.com/pressly/goose).

For detailed database information, see [docs/DATABASE.md](docs/DATABASE.md).

### Core Tables

- **Family**: Owns lists, configured via config files
- **List**: Named lists within a family (grocery, hardware, etc.)
- **Item**: Items that can be added to lists (global or list-specific)

## API Endpoints

### Family Management

- `GET /api/v1/family_name` - Get all lists for a family
- `POST /api/v1/family_name` - Create a new list

### List Management

- `GET /api/v1/family_name/list_name` - Get all items on a list
- `DELETE /api/v1/family_name/list_name` - Delete a list
- `GET /api/v1/family_name/list_name/items` - Get available items for a list

### Item Management

- `POST /api/v1/family_name/list_name/item_or_uuid` - Add item to list
- `DELETE /api/v1/family_name/list_name/uuid` - Remove item from list

## Development

### Prerequisites

- Go 1.21+
- SQLite3
- Make

### Building

```bash
make build
```

### Database Operations

```bash
# Run migrations
make db-up

# Rollback migrations
make db-down

# Create new migration
make db-create NAME=your_migration_name
```

### Testing

```bash
make test
```

## Design Philosophy

This project follows "Grug's Simple Way" - a design philosophy that prioritizes simplicity over features:

- **One screen does everything** - No complex navigation
- **Big buttons** - Easy to use on any device
- **Works like paper** - Familiar mental model
- **No complexity demons** - Resist feature creep

### Success Metrics

- 5-year-old can use without help
- Grandma can use without tech support
- Works on slow internet
- Family doesn't need instructions

## Contributing

See [docs/CONTRIBUTE.md](docs/CONTRIBUTE.md) for contribution guidelines.

## License

[Add your license here]

## Support

This is a self-hosted solution designed for technical families. Support is community-based through issues and pull requests.
