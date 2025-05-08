# Stage 1: Build
FROM golang:1.24.3-alpine AS builder

# Enable Go modules
ENV GO111MODULE=on

# Install git and gcc (needed by some Go libs like gorm/sqlite)
RUN apk update && apk add --no-cache git gcc musl-dev

# Set the working directory inside container
WORKDIR /app

# Copy go mod and sum files first for better caching
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Run
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .

# Expose the Fiber app port
EXPOSE 3000

# Run the binary
CMD ["./main"]
