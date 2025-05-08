# Book API (Golang + Fiber + GORM + SQLite)

This is a simple RESTful Book Management API built using **Go (v1.24.2)**, **Fiber v2**, **GORM**, and **SQLite**. It allows you to create, read, update, delete, search, and paginate book records. The application also includes unit testing and Docker support.

---

## Features

- Create a new book
- Get all books
- Get a single book by ID
- Update a book
- Delete a book
- Search books by title or author
- Paginate book results
- RESTful JSON responses
- Unit tests (using `testify`)
- Docker & Docker Compose support

---

## Technologies Used

- [Go 1.24.2](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [GORM ORM](https://gorm.io/)
- [SQLite (in-memory for tests)](https://www.sqlite.org/)
- [Testify](https://github.com/stretchr/testify)
- [Docker](https://www.docker.com/)

---

## Project Structure

book-api/
│
├── handlers/ # Fiber route handlers
├── services/ # CRUD, Search, Pagination
├── models/ # GORM model definitions
├── database/ # Database connection setup
│
├── main.go # Entry point
├── go.mod # Go module definitions
├── Dockerfile # Docker image config
├── docker-compose.yml # Docker Compose service
└── README.md


---

### Base URL: http://localhost:3000

## 🧪 Running Locally

### 1. Extract and Run with Go

1. 📦 Extract the ZIP file.
2. 📁 Open a terminal inside the extracted folder (`book-api`).
3. Run the following commands:

```bash
go mod tidy
go run main.go

================= OR ======================

```bash
git clone https://github.com/mfarthas-al/book-api.git
cd book-api
go mod tidy
go run main.go


visit:
http://localhost:3000/books


### 2. Run with Docker

docker compose build --no-cache
docker compose up 

Visit:
http://localhost:3000/books


### 3. API End Points

| Method | Endpoint                | Description                  |
| ------ | ----------------------- | ---------------------------- |
| GET    | `/books`                | Get all books                |
| GET    | `/books/:id`            | Get a book by ID             |
| POST   | `/books`                | Create a new book            |
| PUT    | `/books/:id`            | Update a book                |
| DELETE | `/books/:id`            | Delete a book                |
| GET    | `/books?limit=3&page=1` | Paginate books               |
| GET    | `/books/search?q=go`    | Search books by title/author |


### 4. Run Unit Tests

go test ./...



