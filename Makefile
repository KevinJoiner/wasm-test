.PHONY: clean run build install dep test lint format docker

PATHINSTBIN = $(abspath ./bin)
export PATH := $(PATHINSTBIN):$(PATH)
SHELL := env PATH=$(PATH) $(SHELL)

BIN_NAME					?= wasm-app
DEFAULT_INSTALL_DIR			:= $(go env GOPATH)/$(PATHINSTBIN)
DEFAULT_ARCH				:= $(shell go env GOARCH)
DEFAULT_GOOS				:= $(shell go env GOOS)
ARCH						?= $(DEFAULT_ARCH)
GOOS						?= $(DEFAULT_GOOS)
INSTALL_DIR					?= $(DEFAULT_INSTALL_DIR)
.DEFAULT_GOAL 				:= run


VERSION   := $(shell git describe --tags || echo "v0.0.0")
VER_CUT   := $(shell echo $(VERSION) | cut -c2-)

# Dependency versions
GOLANGCI_VERSION   = latest


help:
	@echo "\nSpecify a subcommand:\n"
	@grep -hE '^[0-9a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[0;36m%-20s\033[m %s\n", $$1, $$2}'
	@echo ""

build: ## Build the binary
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(ARCH) \
		go build -o $(PATHINSTBIN)/$(BIN_NAME) ./cmd/$(BIN_NAME)

build-plugin: ## Build WASM plugin
	tinygo build -o $(PATHINSTBIN)/$(BIN_NAME)-plugin.wasm -target wasi ./cmd/$(BIN_NAME)-plugin

run: build-plugin build ## Run the binary
	@$(PATHINSTBIN)/$(BIN_NAME)
all: clean target

clean: ## Clean all build artifacts
	@rm -rf $(PATHINSTBIN)
	

tidy: ## Tidy go modules
	@go mod tidy

test: ## Run tests
	@go test ./...

lint: ## Run linter
	echo "Warning linting does not work currently for wasi files"
	@golangci-lint run

tools-golangci-lint: ## Install golangci-lint
	@mkdir -p $(PATHINSTBIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | BINARY=golangci-lint bash -s -- ${GOLANGCI_VERSION}

tools: tools-golangci-lint  ## Install all tools

generate: go-generate ## run all file generation for the project

go-generate:## run go generate
	@go generate ./...