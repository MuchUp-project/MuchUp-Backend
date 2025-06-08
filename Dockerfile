# --- 1. Build Stage ---
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy dependency files and download them
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# CGO_ENABLED=0 is important for creating a static binary
# -o /app/server creates the binary named 'server' in the /app directory
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server.go


# --- 2. Final Stage ---
FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/server .

# Expose ports for REST API and gRPC
EXPOSE 8080
EXPOSE 50051

# The command to run the application
CMD ["/app/server"] 