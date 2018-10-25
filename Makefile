BINDATA=pkg/manifests/bindata.go
IMAGE_REPOSITORY_NAME ?= openshift

.PHONY: all
all: generate build

.PHONY: deps
deps:
	go get -u github.com/jteeuwen/go-bindata/...

.PHONY: generate
generate: deps
	go-bindata -mode 420 -modtime 1 -pkg manifests -o $(BINDATA) assets/... manifests/...

.PHONY: build
build: generate
	GOOS=$(GOOS) go build -o pod-checkpointer-operator ./cmd/pod-checkpointer-operator

.PHONY: images
images:
	imagebuilder -f Dockerfile -t $(IMAGE_REPOSITORY_NAME)/origin-pod-checkpointer-operator .
