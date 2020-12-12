SHELL = /bin/bash

load-env:
	source ./.env

test: load-env
	go test -race ./...

build-statics:
	pkger -o templates