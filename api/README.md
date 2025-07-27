# API

This is a template API implementation using Go and the Gin framework, with automatic Swagger documentation generation.

## Technology Stack

- **[Go](https://go.dev/)** - Programming language
- **[Gin](https://gin-gonic.com/)** - Web framework
- **[gin-swagger](https://github.com/swaggo/gin-swagger)** - Swagger/OpenAPI documentation generation

## Features

- RESTful API endpoints
- Automatic Swagger/OpenAPI documentation
- **Pure Go OpenAPI 3.0.3 generation** (no npm dependencies)
- CORS support
- Health check endpoint
- Example integration with external API (JSONPlaceholder)
- Proper error handling and response formatting

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Internet connection (for external API calls)

### Installation

1. Navigate to the API directory:
   ```bash
   cd api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Install swag tool (for Swagger generation):
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

4. Ensure swag is in your PATH (add this to your shell profile if needed):
   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

### Running the API

1. Generate Swagger 2.0 documentation:
   ```bash
   # Make sure swag is in PATH
   export PATH=$PATH:$(go env GOPATH)/bin
   swag init
   ```

2. Start the server:
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`

### API Endpoints

- `GET /api/v1/health` - Health check
- `GET /api/v1/posts` - Get all posts from JSONPlaceholder
- `GET /api/v1/posts/{id}` - Get a specific post by ID
- `GET /swagger/index.html` - Swagger UI documentation

### Swagger Documentation

The API automatically generates OpenAPI/Swagger documentation which is:

- Served at `/swagger/index.html` when the server is running
- Generated as JSON file in `docs/swagger.json`
- Copied to `../openapi-specifications/api.swagger.json`
- **Available in OpenAPI 3.0.3 format using pure Go packages** (no npm required)

### Development

To regenerate OpenAPI 3.0.3 documentation after making changes:

**Using Go-only packages (Recommended):**
```bash
./generate-openapi-go.sh
```

**Using npm packages:**
```bash
./generate-openapi.sh
```

**Manual Swagger 2.0 generation:**
```bash
swag init
cp docs/swagger.json ../openapi-specifications/api.swagger.json
```
