FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-service ./cmd/server.go

FROM alpine:latest

COPY --from=builder /app/auth-service /usr/local/bin/auth-service

EXPOSE 8080

CMD ["auth-service"]