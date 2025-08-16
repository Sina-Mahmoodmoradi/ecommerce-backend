# --- Build Stage ---
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# Install dependencies for build
RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ecommerce-backend ./cmd/api

# --- Run Stage ---
FROM alpine:3.19

WORKDIR /app

# Install CA certificates (for HTTPS)
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/ecommerce-backend .
COPY config.yaml .

EXPOSE 8080

CMD ["./ecommerce-backend"]
