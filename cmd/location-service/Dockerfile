# Start with a build stage using a smaller base image
FROM golang:alpine AS builder

# Install git, required for fetching Go dependencies.
# Alpine imags don't include git by default.
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy the source from the specific service directory to the Working Directory inside the container
COPY ./cmd/location-service ./cmd/location-service
COPY ./internal ./internal
COPY ./pkg ./pkg

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# Build the Go app
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o location-service ./cmd/location-service

# Start a new stage from scratch for the final image
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/location-service .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./location-service"]
