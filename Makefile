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


.PHONY: proto
proto:
	rm -rf gen/python || true
	rm -rf gen/pre-python || true
	buf generate protos
	./scripts/generatePipModule.sh

