FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apk add --no-cache gcc musl-dev && \
    CGO_ENABLED=1 GOOS=linux go build -o go-bot ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-bot .

CMD ["./go-bot"]