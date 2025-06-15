FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod and go.sum files first for better caching
COPY gosrc/go.mod gosrc/go.sum* ./


