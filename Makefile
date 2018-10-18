BINDATA=pkg/manifests/bindata.go

.PHONY: all
all: generate build

.PHONY: generate
generate:
	go-bindata -mode 420 -modtime 1 -pkg manifests -o $(BINDATA) assets/... manifests/...

.PHONY: build
build:
	GOOS=$(GOOS) go build -o pod-checkpointer-operator ./cmd/pod-checkpointer-operator
