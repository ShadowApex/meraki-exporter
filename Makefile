IMAGE_NAME := ghcr.io/shadowapex/meraki-exporter
IMAGE_TAG ?= latest
BUILDER_NAME ?= golang:1.19-buster

GO_FILES := $(shell find ./ -name '*.go')
SRC_FILES := $(GO_FILES)

UID := $(shell id -u)
GID := $(shell id -g)

# Builds the exporter binary
meraki-exporter: $(SRC_FILES)
	go build .

# Builds the docker image
image:
	@echo "Building $(IMAGE_NAME):$(IMAGE_TAG)..."
	docker build -f Dockerfile -t $(IMAGE_NAME):$(IMAGE_TAG) .
ifdef GITHUB_TOKEN
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest
endif

# Pushes the docker image to our container repository
.PHONY: push
push: login
	@echo "Pushing image to $(IMAGE_NAME):$(IMAGE_TAG)"
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
ifdef GITHUB_TOKEN
	docker push $(IMAGE_NAME):latest
endif

.PHONY: login
login:
	@echo $(GITHUB_TOKEN) | docker login ghcr.io -u shadowapex --password-stdin

# Publishes a release of the exporter
.PHONY: release
release: image push

# Runs the exporter
run: meraki-exporter
	./meraki-exporter

# Updates dependencies
dep:
	go mod vendor
	go mod tidy
