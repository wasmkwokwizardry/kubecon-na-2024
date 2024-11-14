.PHONY: lint
lint: ## Lint code using golangci-lint
	golangci-lint run --config .golangci.yaml

.PHONY: lint-fix
lint-fix: ## Lint and fix code using golangci-lint
	golangci-lint run --config .golangci.yaml --fix

.PHONY: build
build: static/regex_v1.wasm static/regex_v2.wasm ## Build the wasm binary using tinygo

static/regex_v1.wasm: v1/main.go v1/plugin/plugin.go go.mod go.sum
	tinygo build -o static/regex_v1.wasm -scheduler=none --no-debug -target=wasi ./v1

static/regex_v2.wasm: v2/main.go v2/plugin/plugin.go go.mod go.sum
	tinygo build -o static/regex_v2.wasm -gc=custom -tags=custommalloc -scheduler=none --no-debug -target=wasi ./v2