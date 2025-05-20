# syntax=docker/dockerfile:1
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gin-init main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gin-init .
COPY config ./config
COPY public ./public
COPY logs ./logs
EXPOSE 40020
ENV GIN_MODE=release
CMD ["./gin-init"] 