.PHONY: lint
lint: ## Lint code using golangci-lint
	golangci-lint run --config .golangci.yaml

.PHONY: lint-fix
lint-fix: ## Lint and fix code using golangci-lint
	golangci-lint run --config .golangci.yaml --fix

.PHONY: build
build: static/regex.wasm ## Build the wasm binary using tinygo

static/regex.wasm: cmd/regex/main.go pkg/plugin/regex/regex.go go.mod go.sum
	tinygo build -o static/regex.wasm -scheduler=none --no-debug -target=wasi ./cmd/regex