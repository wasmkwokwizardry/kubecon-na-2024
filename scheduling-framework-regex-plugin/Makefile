GO_VERSION := $(shell awk '/^go /{print $$2}' go.mod|head -n1)

# Manage platform and builders
PLATFORMS ?= linux/amd64,linux/arm64,linux/s390x,linux/ppc64le
BUILDER ?= docker
ifeq ($(BUILDER),podman)
	ALL_FLAG=--all
else
	ALL_FLAG=
endif

# REGISTRY is the container registry to push
# into. The default is to push to the staging
# registry, not production(registry.k8s.io).
REGISTRY?=gcr.io/k8s-staging-scheduler-plugins
RELEASE_VERSION?=v$(shell date +%Y%m%d)-0.0.1
RELEASE_IMAGE:=kube-scheduler:$(RELEASE_VERSION)
GO_BASE_IMAGE?=golang:$(GO_VERSION)
DISTROLESS_BASE_IMAGE?=gcr.io/distroless/static:nonroot
EXTRA_ARGS=""

# VERSION is the scheduler's version
#
# The RELEASE_VERSION variable can have one of two formats:
# v20201009-v0.18.800-46-g939c1c0 - automated build for a commit(not a tag) and also a local build
# v20200521-v0.18.800             - automated build for a tag
VERSION=$(shell echo $(RELEASE_VERSION) | awk -F - '{print $$2}')
VERSION:=$(or $(VERSION),v0.0.$(shell date +%Y%m%d))

# PLUGIN_VERSION is the version of the plugin
PLUGIN_VERSION?=v1

.PHONY: all
all: build

.PHONY: build
build: build-scheduler

.PHONY: build-scheduler
build-scheduler:
	$(GO_BUILD_ENV) go build -ldflags '-X k8s.io/component-base/version.gitVersion=$(VERSION) -w' -o bin/kube-scheduler $(PLUGIN_VERSION)/main.go

.PHONY: build-images
build-images:
	BUILDER=$(BUILDER) \
	PLATFORMS=$(PLATFORMS) \
	RELEASE_VERSION=$(RELEASE_VERSION) \
	REGISTRY=$(REGISTRY) \
	IMAGE=$(RELEASE_IMAGE) \
	GO_BASE_IMAGE=$(GO_BASE_IMAGE) \
	DISTROLESS_BASE_IMAGE=$(DISTROLESS_BASE_IMAGE) \
	DOCKER_BUILDX_CMD=$(DOCKER_BUILDX_CMD) \
	EXTRA_ARGS=$(EXTRA_ARGS) hack/build-images.sh

.PHONY: local-image
local-image: PLATFORMS="linux/$$(uname -m)"
local-image: RELEASE_VERSION="v0.0.0"
local-image: REGISTRY="localhost:5000/scheduler-plugins"
local-image: EXTRA_ARGS="--load"
local-image: clean build-images

.PHONY: release-images
push-images: EXTRA_ARGS="--push"
push-images: build-images

.PHONY: clean
clean: ## Clean up the build directory
	rm -rf ./bin

.PHONY: lint
lint: ## Lint code using golangci-lint
	golangci-lint run --config .golangci.yaml

.PHONY: lint-fix
lint-fix: ## Lint and fix code using golangci-lint
	golangci-lint run --config .golangci.yaml --fix