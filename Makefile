SHELL = /bin/bash

VERSION ?= "v0.1.1"
TAGS ?= ""
GO_BIN ?= "go"

load-env:
	source ./.env

deps:
	$(GO_BIN) get -tags ${TAGS} -t ./...

test: load-env
	go test -race ./...

build-statics:
	pkger -o templates

build: build-statics
	$(GO_BIN) build -v .
lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run --enable-all

install:
	$(GO_BIN) install -tags ${TAGS} -v .

update:
	rm go.*
	$(GO_BIN) mod init
	$(GO_BIN) mod tidy
	make test
	make install

prepare-release:
	git tag -a ${VERSION}
	git push origin ${VERSION}
release-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...


release: build-statics prepare-release
	export GITHUB_TOKEN=$(cat ~/.config/goreleaser/github_token)
	$(GO_BIN) get github.com/gobuffalo/release
	${VERSION} | release -y -f version.go --skip-packr