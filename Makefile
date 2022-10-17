DOCKER_NS ?= github.com/Ning-Qing
PKG_NAME ?= temple
PKG_VERSION ?= 1.13
GO_VERSION ?= 1.17.13
ALPINE_VERSION ?= 3.16
SUPPORTED_PLATFORMS = linux/arm64,linux/amd64


ARCH := $(shell arch)
OS := $(shell uname -s)

ifeq ($(ARCH),x86_64)
	ARCH=amd64
else ifeq ($(ARCH),arm64)
	ARCH=arm64
else 
	exit $$?
endif


ifeq ($(OS),Linux)
	OS=linux
else
	exit $$?
endif


build:
	@echo "Building $(PKG_NAME)..."
	@go mod tidy && go build -o ./build/$(PKG_NAME) \
		-ldflags "-X main.Version=$(PKG_VERSION)" .
	@echo "see ./build/$(PKG_NAME)"
	
run:
	@go mod tidy && go run .

image:
	@echo "Building $(PKG_NAME) docker image - $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION)"
	@docker build --no-cache -f Dockerfile \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg ALPINE_VERSION=$(ALPINE_VERSION) \
		-t $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION) .
	@docker tag $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION) $(PKG_NAME):latest
	@echo "scan image $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION) or $(PKG_NAME):latest"
	@echo "maybe you want psuh image to $(DOCKER_NS),you can make push"

push:
	@echo "push image $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION) to $(DOCKER_NS)"
	@docker push $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION)

image-relase:
	@docker buildx create --use --name $(PKG_NAME)-builder
	@echo "Building $(PKG_NAME) docker image relase - $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION)"
	@docker buildx build \
		--platform $(SUPPORTED_PLATFORMS) \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg ALPINE_VERSION=$(ALPINE_VERSION) \
		-f Dockerfile -t $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION) . --push
	@docker tag $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION) $(PKG_NAME):latest
	@echo "you can go to $(DOCKER_NS) see image $(DOCKER_NS)/$(PKG_NAME):$(PKG_VERSION)"
	@docker buildx rm $(PKG_NAME)-builder

clean:
	@rm -rf build

.PHONY: build run image image-relase
