BINDATA=pkg/manifests/bindata.go
IMAGE_REPOSITORY_NAME ?= openshift

.PHONY: all
all: generate build

.PHONY: generate
generate:
	go-bindata -mode 420 -modtime 1 -pkg manifests -o $(BINDATA) assets/... manifests/...

.PHONY: build
build:
	GOOS=$(GOOS) go build -o pod-checkpointer-operator ./cmd/pod-checkpointer-operator

.PHONY: images
images:
	imagebuilder -f Dockerfile -t $(IMAGE_REPOSITORY_NAME)/origin-pod-checkpointer-operator .
