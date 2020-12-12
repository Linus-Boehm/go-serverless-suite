SHELL = /bin/bash

load-env:
	source ./.env

test: load-env

	go test -race ./...