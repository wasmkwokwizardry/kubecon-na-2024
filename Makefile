.PHONY: start-simulator-wasm-extension
start-simulator-wasm-extension: ## Start the Kubernetes Scheduler Simulator and the WASM Extension
	cd wasm-extension-regex-plugin && make build
	cd kube-scheduler-simulator && make docker_build
	docker compose \
		-f kube-scheduler-simulator/compose.yml 	    \
		-f kube-scheduler-simulator/compose.local.yml   \
		-f wasm-extension-regex-plugin/compose.wasm.yml \
		up -d

.PHONY: stop-simulator-wasm-extension
stop-simulator-wasm-extension: ## Stop the Kubernetes Scheduler Simulator and the WASM Extension
	docker compose \
		-f kube-scheduler-simulator/compose.yml 			\
		-f kube-scheduler-simulator/compose.local.yml 		\
		-f wasm-extension-regex-plugin/compose.wasm.yml \
		down --volumes

.PHONY: start-simulator-regex-extender
start-simulator-regex-extender: ## Start the Kubernetes Scheduler Simulator and the Regex Extender
	cd kube-scheduler-simulator && make docker_build
	docker compose \
		-f kube-scheduler-simulator/compose.yml 	  \
		-f kube-scheduler-simulator/compose.local.yml \
		-f regex-extender/compose.extender.yml 		  \
		up -d

.PHONY: stop-simulator-regex-extender
stop-simulator-regex-extender: ## Stop the Kubernetes Scheduler Simulator and the Regex Extender
	docker compose \
		-f kube-scheduler-simulator/compose.yml 	  \
		-f kube-scheduler-simulator/compose.local.yml \
		-f regex-extender/compose.extender.yml 		  \
		down --volumes

.PHONY: start-simulator
start-simulator: ## Start the Kubernetes Scheduler Simulator
	cd kube-scheduler-simulator && make docker_build_and_up

.PHONY: stop-simulator
stop-simulator: ## Stop the Kubernetes Scheduler Simulator
	cd kube-scheduler-simulator && make docker_down_local