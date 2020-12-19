SHELL = /bin/bash

VERSION ?= "v0.1.9"
TAGS ?= ""
GO_BIN ?= "go1.16beta1"

deps:
	$(GO_BIN) get -tags ${TAGS} -t ./...

test:
	$(GO_BIN) test -race ./...

build-statics:
	pkger -o templates

build: build-statics
	$(GO_BIN) build -v .

lint:
	$(GO_BIN) get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run

install:
	echo "Skip install for this package"

update:
	rm go.*
	$(GO_BIN) mod init
	$(GO_BIN) mod tidy
	make test
	make install

release-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...

release: release-test build-statics
	$(GO_BIN) get github.com/gobuffalo/release
	release -y -f version.go -v ${VERSION} --skip-packr