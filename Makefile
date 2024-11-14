compose := cd kube-scheduler-simulator && docker compose -f compose.yml -f compose.local.yml
compose_up := up -d --remove-orphans
compose_down := down --volumes --remove-orphans

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display the help for supported commands.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Kube Scheduler WASM Extension

.PHONY: start-simulator-wasm
start-simulator-wasm: wasm-build simulator-build ## Start the Kubernetes Scheduler Simulator and the WASM Extension
	@$(compose) -f ../wasm-extension-regex-plugin/compose.wasm.yml $(compose_up)

.PHONY: wasm-build
wasm-build: ## Build the WASM Extension
	@cd wasm-extension-regex-plugin && make build

##@ Kube Scheduler Extender

.PHONY: start-simulator-extender
start-simulator-extender: simulator-build ## Start the Kubernetes Scheduler Simulator and the Regex Extender
	@$(compose) -f ../scheduler-extender-regex/compose.extender.yml $(compose_up) --build

.PHONY: start-simulator-extender-v2
start-simulator-extender-v2: simulator-build ## Start the Kubernetes Scheduler Simulator and the Regex Extender
	@$(compose) -f ../scheduler-extender-regex/compose.extender.yml -f ../scheduler-extender-regex/compose.extender.v2.yml $(compose_up) --build

##@ Helpers

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
simulator-submodule: kube-scheduler-simulator/.git ## Initialize the git kube-scheduler-simulator submodule

kube-scheduler-simulator/.git:
	@git submodule update --init --recursive kube-scheduler-simulator