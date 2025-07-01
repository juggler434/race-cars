# Race Cars API

A RESTful API for managing race cars built with Go and Gorilla Mux.

## Features

- **RESTful API**: Full CRUD operations for race cars
- **PostgreSQL Database**: Robust data storage with migrations
- **Gorilla Mux Router**: Fast and flexible HTTP routing
- **Middleware Support**: CORS, logging, and error recovery
- **Environment Configuration**: Flexible configuration management
- **JSON API**: Standard JSON request/response format

## Project Structure

```
race-cars/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── env.example             # Environment variables template
├── internal/
│   ├── config/             # Configuration management
│   │   └── config.go
│   ├── models/             # Data models
│   │   └── car.go
│   ├── repository/         # Database operations
│   │   └── car_repository.go
│   ├── handlers/           # HTTP request handlers
│   │   └── car_handler.go
│   ├── middleware/         # HTTP middleware
│   │   └── middleware.go
│   └── routes/             # Route definitions
│       └── routes.go
└── migrations/             # Database migrations
    └── 001_create_cars_table.sql
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher

## Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd race-cars
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up the database**
   ```bash
   # Create PostgreSQL database
   createdb race_cars
   
   # Run migrations
   psql -d race_cars -f migrations/001_create_cars_table.sql
   ```

4. **Configure environment variables**
   ```bash
   cp env.example .env
   # Edit .env with your database credentials
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`

## API Endpoints

### Cars

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/cars` | Get all cars |
| GET | `/api/cars/{id}` | Get car by ID |
| POST | `/api/cars` | Create a new car |
| PUT | `/api/cars/{id}` | Update a car |
| DELETE | `/api/cars/{id}` | Delete a car |

### Other Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | API information |
| GET | `/health` | Health check |

## Request/Response Examples

### Create a Car

**Request:**
```bash
POST /api/cars
Content-Type: application/json

{
  "name": "Ferrari F40",
  "brand": "Ferrari",
  "model": "F40",
  "year": 1987,
  "engine_size": 2.9,
  "horsepower": 471,
  "top_speed": 324,
  "weight": 1100,
  "category": "Supercar",
  "description": "The Ferrari F40 is a mid-engine, rear-wheel drive sports car.",
  "image_url": "https://example.com/ferrari-f40.jpg"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Ferrari F40",
    "brand": "Ferrari",
    "model": "F40",
    "year": 1987,
    "engine_size": 2.9,
    "horsepower": 471,
    "top_speed": 324,
    "weight": 1100,
    "category": "Supercar",
    "description": "The Ferrari F40 is a mid-engine, rear-wheel drive sports car.",
    "image_url": "https://example.com/ferrari-f40.jpg",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

### Get All Cars

**Request:**
```bash
GET /api/cars
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Ferrari F40",
      "brand": "Ferrari",
      "model": "F40",
      "year": 1987,
      "engine_size": 2.9,
      "horsepower": 471,
      "top_speed": 324,
      "weight": 1100,
      "category": "Supercar",
      "description": "The Ferrari F40 is a mid-engine, rear-wheel drive sports car.",
      "image_url": "https://example.com/ferrari-f40.jpg",
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  ]
}
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | `password` |
| `DB_NAME` | Database name | `race_cars` |
| `DB_SSLMODE` | Database SSL mode | `disable` |
| `DATABASE_URL` | Full database URL (alternative) | - |
| `PORT` | Server port | `8080` |
| `LOG_LEVEL` | Logging level | `info` |

## Development

### Running Tests
```bash
go test ./...
```

### Building
```bash
go build -o race-cars-api main.go
```

### Running with Docker
```bash
# Build Docker image
docker build -t race-cars-api .

# Run container
docker run -p 8080:8080 --env-file .env race-cars-api
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.
