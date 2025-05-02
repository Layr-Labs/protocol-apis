.PHONY: build clean test

GO = $(shell which go)
BIN = ./bin

PROTO_OPTS=--proto_path=protos --go_out=paths=source_relative:protos

all: build

.PHONY: deps/system
deps/system:
	./scripts/installDeps.sh

.PHONY: deps/local-plugin
deps/local-plugin:
	git submodule update --init --recursive
	cd buf-plugin-openapi && make clean install


.PHONY: deps/dev
deps/dev: deps/system deps/local-plugin
	${GO} get google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	${GO} get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	${GO} get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	${GO} get github.com/grpc-ecosystem/protoc-gen-grpc-gateway-ts
	${GO} get github.com/akuity/grpc-gateway-client/protoc-gen-grpc-gateway-client@latest
	buf dep update

.PHONY: deps/go
deps/go:
	${GO} mod tidy
	npm install

.PHONY: proto
proto:
	rm -rf gen/python || true
	rm -rf gen/pre-python || true
	buf generate
	# ./scripts/generatePipModule.sh
	# mkdir gen/openapiv3 || true
	# ./node_modules/.bin/swagger2openapi --outfile gen/openapiv3/openapi.yaml gen/openapiv2/apidocs.swagger.json

