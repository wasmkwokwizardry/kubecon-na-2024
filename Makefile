.PHONY: start-simulator-wasm-extension
start-simulator-wasm-extension: submodules ## Start the Kubernetes Scheduler Simulator and the WASM Extension
	@cd wasm-extension-regex-plugin && make build
	@cd kube-scheduler-simulator && make docker_build
	@docker compose										\
		-p kube-scheduler-simulator						\
		-f kube-scheduler-simulator/compose.yml			\
		-f kube-scheduler-simulator/compose.local.yml	\
		-f wasm-extension-regex-plugin/compose.wasm.yml	\
		up -d --remove-orphans

.PHONY: stop-simulator-wasm-extension
stop-simulator-wasm-extension: submodules ## Stop the Kubernetes Scheduler Simulator and the WASM Extension
	@docker compose										\
		-p kube-scheduler-simulator						\
		-f kube-scheduler-simulator/compose.yml			\
		-f kube-scheduler-simulator/compose.local.yml	\
		-f wasm-extension-regex-plugin/compose.wasm.yml	\
		down --volumes --remove-orphans

.PHONY: start-simulator-regex-extender
start-simulator-regex-extender: submodules ## Start the Kubernetes Scheduler Simulator and the Regex Extender
	@cd kube-scheduler-simulator && make docker_build
	@docker compose										\
		-p kube-scheduler-simulator						\
		-f kube-scheduler-simulator/compose.yml			\
		-f kube-scheduler-simulator/compose.local.yml	\
		-f regex-extender/compose.extender.yml			\
		up -d --remove-orphans --build

.PHONY: stop-simulator-regex-extender
stop-simulator-regex-extender: submodules ## Stop the Kubernetes Scheduler Simulator and the Regex Extender
	@docker compose										\
		-p kube-scheduler-simulator						\
		-f kube-scheduler-simulator/compose.yml			\
		-f kube-scheduler-simulator/compose.local.yml	\
		-f regex-extender/compose.extender.yml			\
		down --volumes --remove-orphans

.PHONY: start-simulator
start-simulator: submodules ## Start the Kubernetes Scheduler Simulator
	@cd kube-scheduler-simulator && make docker_build
	@docker compose										\
		-p kube-scheduler-simulator						\
		-f kube-scheduler-simulator/compose.yml			\
		-f kube-scheduler-simulator/compose.local.yml	\
		up -d --remove-orphans

.PHONY: stop-simulator
stop-simulator: submodules ## Stop the Kubernetes Scheduler Simulator
	@docker compose										\
		-p kube-scheduler-simulator						\
		-f kube-scheduler-simulator/compose.yml			\
		-f kube-scheduler-simulator/compose.local.yml	\
		down --volumes --remove-orphans

.PHONY: submodules
submodules: kube-scheduler-simulator/.git ## Initialize the git submodules

kube-scheduler-simulator/.git:
	@git submodule update --init --recursive kube-scheduler-simulator