FROM golang:1.24.6 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

FROM alpine:3.18
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/server .
COPY .env .

EXPOSE 8081

CMD ["./server"]