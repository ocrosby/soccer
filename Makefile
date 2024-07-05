.PHONY: deps, lint

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

# Target for running linter
lint:
	@golangci-lint run ./pkg/tgs/...
