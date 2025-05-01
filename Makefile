.PHONY: build clean test

GO = $(shell which go)
BIN = ./bin

PROTO_OPTS=--proto_path=protos --go_out=paths=source_relative:protos

all: build

.PHONY: deps
deps: deps/dev deps/go
	./scripts/installDeps.sh
	${GO} install github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts@latest
	cd protos && buf dep update

.PHONY: deps/dev
deps/dev: pre-build

.PHONY: deps/go
deps/go:
	${GO} mod tidy
	npm install

pre-build:
	git submodule update --init --recursive
	cd buf-plugin-openapi && make clean install

.PHONY: proto
proto: pre-build
	rm -rf gen/python || true
	rm -rf gen/pre-python || true
	buf generate protos
	# ./scripts/generatePipModule.sh
	# mkdir gen/openapiv3 || true
	# ./node_modules/.bin/swagger2openapi --outfile gen/openapiv3/openapi.yaml gen/openapiv2/apidocs.swagger.json

