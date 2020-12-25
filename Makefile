SHELL = /bin/bash

VERSION ?= "v0.6.0"
TAGS ?= ""
GO_BIN ?= "go1.16beta1"

dynamodb-local:
	cd dockerfiles/dynamodb && docker-compose up -d --no-recreate

deps:
	$(GO_BIN) get -tags ${TAGS} -t ./...

test:
	$(GO_BIN) test -race ./...

integration-test
	INTEGRATIONTEST_ON=true $(GO_BIN) test -race ./...

build:
	$(GO_BIN) build -v .

lint:
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

release: release-test
	$(GO_BIN) get github.com/gobuffalo/release
	release -y -f version.go -v ${VERSION} --skip-packr