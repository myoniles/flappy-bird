# ---- Config -----------------------------------------------------------------
APP      := vulcan
PKG      := ./cmd/$(APP)
BIN_DIR  := bin
WEB_DIR  := web
PORT     ?= 8080

# wasm_exec.js moved from misc/wasm to lib/wasm in Go 1.24; check both.
GOROOT_DIR := $(shell go env GOROOT)
WASM_EXEC  := $(firstword $(wildcard $(GOROOT_DIR)/lib/wasm/wasm_exec.js $(GOROOT_DIR)/misc/wasm/wasm_exec.js))

.DEFAULT_GOAL := help
.PHONY: help run build test cover fmt vet check build-web serve-web clean

# ---- Help -------------------------------------------------------------------
help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

# ---- Desktop ----------------------------------------------------------------
run: ## Run the game natively
	go run $(PKG)

build: ## Build a native binary into bin/
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP) $(PKG)

# ---- Quality ----------------------------------------------------------------
test: ## Run all tests
	go test ./...

cover: ## Run tests with a coverage report
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

fmt: ## Format all Go code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

check: fmt vet test ## Format, vet, and test (run before committing)

# ---- Web --------------------------------------------------------------------
build-web: ## Build the WebAssembly version into web/
	@test -n "$(WASM_EXEC)" || (echo "wasm_exec.js not found in GOROOT" && exit 1)
	GOOS=js GOARCH=wasm go build -o $(WEB_DIR)/$(APP).wasm $(PKG)
	cp "$(WASM_EXEC)" $(WEB_DIR)/

serve-web: build-web ## Build and serve the web version on localhost:$(PORT)
	@echo "Serving at http://localhost:$(PORT)"
	python3 -m http.server $(PORT) -d $(WEB_DIR)

# ---- Cleanup ----------------------------------------------------------------
clean: ## Remove build artifacts
	rm -rf $(BIN_DIR) coverage.out
	rm -f $(WEB_DIR)/$(APP).wasm $(WEB_DIR)/wasm_exec.js
