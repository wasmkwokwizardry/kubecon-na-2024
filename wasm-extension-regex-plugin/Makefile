.PHONY: lint
lint: ## Lint code using golangci-lint
	golangci-lint run --config .golangci.yaml

.PHONY: lint-fix
lint-fix: ## Lint and fix code using golangci-lint
	golangci-lint run --config .golangci.yaml --fix

.PHONY: build
build: ## Build the wasm binary using tinygo
	tinygo build -o bin/regex.wasm -scheduler=none --no-debug -target=wasi cmd/regex/main.go