# Go Sample API

A production-oriented REST API sample built with Go and Gin.

This project demonstrates common backend engineering patterns and can be adapted into various architectural styles such like Clean Architecture, Domain Driven Design (DDD), Modular Monolith, etc.

The project intentionally keeps the implementation simple while preserving extensibility and scalability.

---

# Features
- Health Check Endpoint /ping
- Echo Endpoint
- CRUD Book API
- Search & Pagination
- Graceful Shutdown
- Authentication Middleware
- Auth Provider
  - JWT Authentication
  - In-Memory Token
- Repository Pattern
- Swappable Auth Provider Adapter
- Swappable Persistence Layer

---

# Project Goals

This repository is designed as a:

- starter project
- learning reference
- backend template
- architecture playground

The implementation is intentionally flexible and not tied to a single architectural approach.

You can evolve this project into:

- enterprise backend service
- microservice
- monolith
- internal API service
- event-driven service

---


# Extensibility

This project is intentionally designed to be extensible.

The current implementation uses:

- in-memory persistence
- in-memory token authentication

However, the architecture allows these components to be replaced without changing the business logic.

---

# Persistence Extensibility

Current implementation:

```go
bookRepo := book.NewMemoryStore()
```
The repository layer is abstracted through interfaces, allowing the storage implementation to be swapped easily.

## Example PostgreSQL Repository Adapter

```go
bookRepo := postgres.NewBookRepository(db)
```

# Installation

Clone repository:

```bash
git clone https://github.com/example/go-sample-api.git
```

Move into project directory:

```bash
cd go-sample-api
```

Install dependencies:

```bash
go mod tidy
```

---

# Run Application

```bash
go run cmd/api/main.go
```

Server will run on:

```txt
http://localhost:8080
```

---

# Docker Compose

The project includes Docker Compose configuration for:

- Go API
- PostgreSQL
- Redis
- Nginx Reverse Proxy

---

## Start Services

```bash
docker compose up --build
```

Run detached:

```bash
docker compose up -d --build
```

Stop services:

```bash
docker compose down
```

---

# API Endpoints

---

## Health Check

### GET `/ping`

### Request

```bash
curl localhost:8080/ping
```

### Response

```json
{
  "message": "pong"
}
```

---

## Echo Endpoint

### POST `/echo`

### Request

```bash
curl -X POST localhost:8080/echo \
-H "Content-Type: application/json" \
-d '{
  "hello":"world"
}'
```

### Response

```json
{
  "echo": {
    "hello": "world"
  }
}
```

---

# Authentication

Protected endpoints require:

```txt
Authorization: Bearer <token>
```

---

## Default Token

```txt
bypass-token
```

Example:

```bash
-H "Authorization: Bearer bypass-token"
```

---

# Book API

---

## Create Book

### POST `/api/books`

### Request

```bash
curl -X POST localhost:8080/api/books \
-H "Authorization: Bearer bypass-token" \
-H "Content-Type: application/json" \
-d '{
  "title":"Clean Code",
  "author":"Robert Martin"
}'
```

---

## Get Book By ID

### GET `/api/books/:id`

### Request

```bash
curl localhost:8080/api/books/{id} \
-H "Authorization: Bearer bypass-token"
```

---

## Update Book

### PUT `/api/books/:id`

### Request

```bash
curl -X PUT localhost:8080/api/books/{id} \
-H "Authorization: Bearer bypass-token" \
-H "Content-Type: application/json" \
-d '{
  "title":"Clean Architecture",
  "author":"Robert Martin"
}'
```

---

## Delete Book

### DELETE `/api/books/:id`

### Request

```bash
curl -X DELETE localhost:8080/api/books/{id} \
-H "Authorization: Bearer bypass-token"
```

---

## Search Book

### GET `/api/books`

### Query Parameters

| Parameter | Description |
|---|---|
| query | Search by title |
| page | Page number |
| limit | Items per page |

---

### Example

```bash
curl "localhost:8080/api/books?query=clean&page=1&limit=10" \
-H "Authorization: Bearer bypass-token"
```

### Example Response

```json
{
  "success": true,
  "data": {
    "items": [
      {
        "id": "uuid",
        "title": "Clean Code",
        "author": "Robert Martin"
      }
    ],
    "meta": {
      "page": 1,
      "limit": 10,
      "total": 1
    }
  }
}
```

---

# Error Response Format

All API errors use a consistent response format.

Example:

```json
{
  "success": true,
  "data": {},
  "code": "NOT_FOUND",
  "message": "book not found"
}
```

---

# Graceful Shutdown

The application supports graceful shutdown.

Supported signals:

- SIGINT
- SIGTERM

Benefits:

- ongoing requests can finish properly
- safe rolling deployment
- prevents abrupt connection termination
- suitable for Docker and Kubernetes environments

---


# Tech Stack

- Go
- Gin
- JWT
- UUID

---

# Project Structure

```txt
go-sample-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── provider/
│   │   ├── auth.go
│   │   ├── jsonwebtoken.go
│   │   ├── memory-token.go
│   ├── book/
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── service.go
│   ├── httpx/
│   │   ├── handler.go
│   │   └── middleware.go
│   │   └── respond.go
│   └── repository/
│       ├── memory.go
├── go.mod
├── docker-compose.yaml
└── README.md

# License

MIT