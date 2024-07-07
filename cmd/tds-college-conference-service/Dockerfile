# Define variables for the build
ARG GO_VERSION=1.21-alpine
ARG ALPINE_VERSION=3.14
ARG PORT=8080

# Enable Docker BuildKit
# syntax=docker/dockerfile:1.2

# Use the latest stable Go image
FROM golang:${GO_VERSION} AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod .
COPY go.sum .


RUN go mod download && \
    go mod verify

# Copy the entire source directory from both cmd and pkg
COPY cmd/tds-college-conference-service cmd/
COPY pkg/ pkg/

# Build the Go application inside the container
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/tds-college-conference-service -ldflags="-s -w" -v -x ./cmd/tds-college-conference-service/service.go


# Use a minimal Alpine image
FROM alpine:${ALPINE_VERSION} AS deploy

# RUN apk upgrade --no-cache

RUN mkdir /app && chmod 777 /app

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/tds-college-conference-service .

# Make the binary executable
RUN chmod +x /app/tds-college-conference-service

# Expose a port (if your Go application listens on a specific port)
EXPOSE 8080

# Add a health check (update the command based on your app's health check endpoint)
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Define the command to run your application
ENTRYPOINT sh -c "/app/tds-college-conference-service serve"
