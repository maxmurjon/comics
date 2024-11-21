# Build step (builder stage)
FROM golang:1.23.1-alpine AS builder  # Go versiyasini va Alpine tasvirini belgilash

RUN mkdir /app

COPY go.mod go.sum /app/
WORKDIR /app

RUN go mod tidy

COPY . /app

RUN go build -o main cmd/main.go

# Final stage (runtime stage)
FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/main /app/main

ENV DOT_ENV_PATH=config/.env

CMD ["/app/main"]
