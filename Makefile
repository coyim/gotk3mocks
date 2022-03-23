GO_FILES := *.go
SOURCE_FILES := $(GO_FILES)

GO := go
GOBUILD := $(GO) build
GOTEST := $(GO) test

.PHONY := default build test

default: test

build:
	$(GOBUILD) ./...

test:
	$(GOTEST) ./...
