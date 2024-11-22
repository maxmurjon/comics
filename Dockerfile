# 1. Asosiy image ni tanlash
FROM golang:1.23.3-alpine AS builder

# 2. Ishchi katalogni yaratish va uni tanlash
WORKDIR /app

# 3. Fayllarni container ichiga ko'chirish
COPY go.mod go.sum ./
RUN go mod download

# 4. Kodni ko'chirish va ilovani qurish
COPY . .
RUN go build -o main cmd/main.go

# 5. Yengil image ni tayyorlash
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main ./

# 6. Muhit o'zgaruvchilarini o'rnatish
ENV DOT_ENV_PATH=config/.env

# 7. Ilovani ishga tushirish
CMD ["/app/main"]
