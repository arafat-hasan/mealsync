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

API documentation will be available at `/swagger/index.html` when the server is running.

## License

MIT 