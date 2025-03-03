.PHONY: build clean test

GO = $(shell which go)
BIN = ./bin

PROTO_OPTS=--proto_path=protos --go_out=paths=source_relative:protos

all: build

deps: deps/go
	./scripts/installDeps.sh

.PHONY: deps/dev
deps/dev:
	${GO} get github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts
	cd protos && buf dep update

.PHONY: deps
deps/go:
	${GO} mod tidy
	npm install

pre-build:
	cd protocol-apis-annotations && make proto
	cd buf-plugin-openapi && make install

.PHONY: proto
proto: pre-build
	rm -rf gen/python || true
	rm -rf gen/pre-python || true
	buf generate protos
	# ./scripts/generatePipModule.sh
	# mkdir gen/openapiv3 || true
	# ./node_modules/.bin/swagger2openapi --outfile gen/openapiv3/openapi.yaml gen/openapiv2/apidocs.swagger.json

