.PHONY: deps, lint, build, run

# Define variables for image name and default tag
IMAGE_NAME=$(SERVICE)-service
TAG=$(shell git rev-parse --short HEAD) # Use git commit hash for uniqueness

# Detect operating system
OS := $(shell uname)

# Target for installing development dependencies
deps:
ifeq ($(OS), Darwin) # macOS
	@which golangci-lint || brew install golangci-lint
else # Linux and others
	@which golangci-lint || { \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.31.0; \
	}
endif

	go install github.com/cucumber/godog/cmd/godog@latest

# Target for running linter
lint:
	@golangci-lint run ./pkg/tgs/...

# Build the Docker image for a specific service
# Usage: make build SERVICE=service-name ENV=dev
build:
	docker build --build-arg ENV=$(ENV) --build-arg VERSION=$(TAG) -t $(SERVICE)-service:$(TAG) -f cmd/$(SERVICE)-service/Dockerfile .

run:
	docker run -p 8080:8080 $(SERVICE)-service:$(TAG)

# Push the Docker image to a registry
push:
	docker push $(IMAGE_NAME):$(TAG)

docker:
	docker build -t tds-coaching-change-service:latest -f services/tds-coaching-change-service/Dockerfile .

spin:
	docker run -d -p 8080:8080 --name tds-coaching-change-service-container tds-coaching-change-service:latest

swagger:
	docker run -d -p 8080:8080 --name tds-coaching-change-service-container tds-coaching-change-service:latest

# Generate swagger documentation
swagger:
	swagger generate spec -o ./services/tds-coaching-change-service/swagger.yaml --scan-models ./internal/tds-coaching-change-service ./services/tds-coaching-change-service

# Example usage: make build ENV=dev
# This allows for dynamic building of images based on the current git commit and specified environment

