compose := cd kube-scheduler-simulator && docker compose -f compose.yml -f compose.local.yml
compose_up := up -d --remove-orphans
compose_down := down --volumes --remove-orphans

.PHONY: start-simulator-wasm
start-simulator-wasm: simulator-build wasm-build ## Start the Kubernetes Scheduler Simulator and the WASM Extension
	@$(compose) -f ../wasm-extension-regex-plugin/compose.wasm.yml $(compose_up)

.PHONY: wasm-build
wasm-build: ## Build the WASM Extension
	@cd wasm-extension-regex-plugin && make build

.PHONY: stop-simulator-wasm
stop-simulator-wasm: simulator-submodule ## Stop the Kubernetes Scheduler Simulator and the WASM Extension
	@$(compose) -f ../wasm-extension-regex-plugin/compose.wasm.yml $(compose_down)

.PHONY: start-simulator-extender
start-simulator-extender: simulator-build ## Start the Kubernetes Scheduler Simulator and the Regex Extender
	@$(compose) -f ../scheduler-extender-regex/compose.extender.yml $(compose_up) --build

.PHONY: stop-simulator-extender
stop-simulator-extender: simulator-submodule ## Stop the Kubernetes Scheduler Simulator and the Regex Extender
	@$(compose) -f ../scheduler-extender-regex/compose.extender.yml $(compose_down)

.PHONY: start-simulator
start-simulator: simulator-build ## Start the Kubernetes Scheduler Simulator
	@$(compose) $(compose_up)

.PHONY: stop-simulator
stop-simulator: simulator-submodule ## Stop the Kubernetes Scheduler Simulator
	@$(compose) $(compose_down)

.PHONY: simulator-build
simulator-build: simulator-submodule ## Build the Kubernetes Scheduler Simulator
	@cd kube-scheduler-simulator && make docker_build

.PHONY: simulator-submodule
simulator-submodule: kube-scheduler-simulator/.git ## Initialize the git simulator-submodule

kube-scheduler-simulator/.git:
	@git submodule update --init --recursive kube-scheduler-simulator