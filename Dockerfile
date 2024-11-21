# Build step (builder stage)
FROM golang:1.23.1-alpine AS builder  # Golang versiyasini Alpine bilan ishlatish

# Application directory in the container
RUN mkdir /app

# Copy the Go modules and Go source code
COPY go.mod go.sum /app/
WORKDIR /app

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . /app

# Build the Go application
RUN go build -o main cmd/main.go

# Final stage (runtime stage)
FROM alpine:3.16

# Set the working directory for the app in the final image
WORKDIR /app

# Copy the compiled binary from the builder image
COPY --from=builder /app/main /app/main

# Set environment variable for configuration
ENV DOT_ENV_PATH=config/.env

# Run the application
CMD ["/app/main"]
