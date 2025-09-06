# syntax=docker/dockerfile:1

FROM golang:1.25 AS builder
WORKDIR /app

# copy mod files first for caching
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the source
COPY . .

# build the binary with proper flags
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/go-starter-template

# runtime container
FROM debian:bullseye-slim
WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server .
RUN chmod +x /app/server

# Expose the port
EXPOSE 8080

CMD ["./server"]
