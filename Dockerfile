FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o api ./cmd/api


FROM alpine:latest

LABEL org.opencontainers.image.source https://github.com/ezflow-me/identity-management-service

ENV SERVER_HOST=0.0.0.0
ENV SERVER_PORT=8080
ENV FIBER_PREFORK=false

WORKDIR /

COPY --from=builder /app/api /api
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/api"]
