FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod and go.sum files first for better caching
COPY go.mod go.sum* ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 go build -o /app/server ./gosrc

# Use a smaller image for the final stage
FROM debian:bullseye-slim

# Install necessary packages for SSL and basic tools
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Copy any required assets or config files
COPY --from=builder /app/shmfile ./shmfile

EXPOSE 5000

CMD ["/app/server"]