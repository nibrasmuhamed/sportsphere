FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN  sh build.sh

FROM alpine:latest

WORKDIR /app

RUN adduser -D appuser
USER appuser

COPY --from=builder /app/sportsphere .
COPY --from=builder /app/config/config.json ./config/config.json

EXPOSE 8000

CMD ["./sportsphere"]