FROM golang:1.19.1-alpine3.16 AS builder

RUN mkdir app

COPY . /app

WORKDIR /app

RUN go build -o main cmd/main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app .
ENV DOT_ENV_PATH=config/.env

CMD ["/app/main"]
