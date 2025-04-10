# MealSync

A Meal Request and Estimation System for managing employee lunch and snacks requests.

## Features

- User Management (Employee & Admin roles)
- Menu Management
- Meal Request System
- Estimation Dashboard
- Attendance Integration
- Notification System

## Tech Stack

- Language: Go
- Web Framework: Gin
- ORM: GORM
- Database: PostgreSQL
- Authentication: JWT

## Project Structure

```
.
├── cmd/                    # Application entry points
├── configs/               # Configuration files
├── docs/                  # Documentation
├── internal/              # Private application code
│   ├── api/              # API handlers
│   ├── middleware/       # HTTP middleware
│   ├── models/          # Data models
│   ├── repository/      # Database operations
│   └── service/         # Business logic
├── pkg/                  # Public libraries
└── scripts/             # Build and deployment scripts
```

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and update the values
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the application:
   ```bash
   go run cmd/main.go
   ```

## API Documentation

The API documentation is available in two formats:

1. **Swagger UI**: Available at `http://localhost:8080/swagger/index.html` when the server is running
2. **Redoc**: A more beautiful and interactive documentation view

### Viewing Documentation

1. Start the server:
   ```bash
   docker-compose up
   ```

2. Access the documentation:
   - Swagger UI: http://localhost:8080/swagger/index.html
   - Redoc: http://localhost:8080/docs

### Documentation Features

- Interactive API testing
- Authentication flow documentation
- Request/Response examples
- Schema definitions
- Error responses
- Security requirements

## License

MIT

## Development Setup with Docker

This project uses Docker for development to ensure a consistent environment across all developers.

### Prerequisites

- Docker
- Docker Compose

### Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/arafat-hasan/mealsync.git
   cd mealsync
   ```

2. Create a `.env` file in the root directory (or copy the example):
   ```
   cp .env.example .env
   ```

3. Start the development environment:
   ```
   docker-compose up
   ```

   This will start:
   - PostgreSQL database on port 5432
   - The MealSync application on port 8080

4. To stop the development environment:
   ```
   docker-compose down
   ```

   To also remove the database volume (this will delete all data):
   ```
   docker-compose down -v
   ```

### Development Workflow

- The application code is mounted as a volume, so changes will be reflected immediately
- The database data is persisted in a Docker volume
- You can access the database using any PostgreSQL client with these credentials:
  - Host: localhost
  - Port: 5432
  - User: postgres
  - Password: postgres
  - Database: mealsync

### API Endpoints

- `POST /api/register` - Register a new user
- `POST /api/login` - Login and get JWT token
- `GET /api/menu` - Get menu items (protected)
- `POST /api/meal-request` - Create a meal request (protected)
- `POST /api/admin/menu` - Create a menu item (admin only)
- `PUT /api/admin/menu/:id` - Update a menu item (admin only)
- `DELETE /api/admin/menu/:id` - Delete a menu item (admin only)
- `GET /api/admin/meal-requests/stats` - Get meal request statistics (admin only) 