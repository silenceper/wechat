export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE := auto
LDFLAGS := -s -w

all: fmt govet golint gotest

gotest:
	go test -v -race ./...

govet:
	go vet ./...

golint:
	golint -set_exit_status $(go list ./...)

fmt:
	go fmt ./...