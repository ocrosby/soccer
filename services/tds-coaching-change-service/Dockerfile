# Start with a build stage using a smaller base image
FROM golang:alpine AS builder

# Install git, required for fetching Go dependencies.
# Alpine images don't include git by default.
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Use the argument to copy the source from the specific service directory to the Working Directory inside the container
COPY ./services/tds-coaching-change-service/swagger.yaml ./swagger.yaml
COPY ./services/tds-coaching-change-service ./services/tds-coaching-change-service
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./shared/swagger-ui ./swagger-ui

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# Use the argument to build the Go app
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tds-coaching-change-service ./services/tds-coaching-change-service

# Start a new stage from scratch for the final image
FROM alpine:latest

WORKDIR /root/

# Use the argument to copy the Pre-built binary file from the previous stage
COPY --from=builder /app/tds-coaching-change-service .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./tds-coaching-change-service"]