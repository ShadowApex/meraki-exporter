IMAGE_NAME := shadowapex/meraki-exporter
IMAGE_TAG ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
BUILDER_NAME ?= golang:1.19-buster

GO_FILES := $(shell find ./ -name '*.go')
SRC_FILES := $(GO_FILES)

UID := $(shell id -u)
GID := $(shell id -g)

meraki-exporter: $(SRC_FILES)
	go build .

image:
	@echo "Building $(IMAGE_NAME):$(IMAGE_TAG)..."
	docker build -f Dockerfile -t $(IMAGE_NAME):$(IMAGE_TAG) .

run: meraki-exporter
	./meraki-exporter

dep:
	go mod vendor
	go mod tidy
