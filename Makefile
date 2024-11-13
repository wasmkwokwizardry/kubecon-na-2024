.PHONY: start-simulator-regex-extender
start-simulator-regex-extender: ## Start the Kubernetes Scheduler Simulator and the Regex Extender
	docker compose \
		-f kube-scheduler-simulator/compose.yml \
		-f kube-scheduler-simulator/compose.yml \
		-f regex-extender/compose.extender.yml \
		up -d

.PHONY: stop-simulator-regex-extender
stop-simulator-regex-extender: ## Stop the Kubernetes Scheduler Simulator and the Regex Extender
	docker compose \
		-f kube-scheduler-simulator/compose.yml \
		-f kube-scheduler-simulator/compose.yml \
		-f regex-extender/compose.extender.yml \
		down --volumes

.PHONY: start-simulator
start-simulator: ## Start the Kubernetes Scheduler Simulator
	cd kube-scheduler-simulator && make docker_up_local

.PHONY: stop-simulator
stop-simulator: ## Stop the Kubernetes Scheduler Simulator
	cd kube-scheduler-simulator && make docker_down_local