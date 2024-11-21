# Build step (builder stage)
FROM golang:1.23.3-alpine AS builder  # Using a valid tag

RUN mkdir /app

# Copy Go modules and Go source code
COPY go.mod go.sum /app/
WORKDIR /app

RUN go mod tidy

# Copy the rest of the application code
COPY . /app

# Build the Go application
RUN go build -o main cmd/main.go

# Final stage (runtime stage)
FROM alpine:3.16

WORKDIR /app

# Copy the compiled binary from the builder image
COPY --from=builder /app/main /app/main

# Set environment variable for configuration
ENV DOT_ENV_PATH=config/.env

# Run the application
CMD ["/app/main"]
