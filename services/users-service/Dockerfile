FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o users-service ./cmd/main.go

FROM alpine:latest

COPY --from=builder /app/users-service /usr/local/bin/users-service

EXPOSE 8080

CMD ["users-service"]