# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go build -o fathers-day .

# Run stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/fathers-day .

EXPOSE 8080

CMD ["./fathers-day"]
