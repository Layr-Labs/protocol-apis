.PHONY: build clean test install generate

# Binary name
BINARY_NAME=protoc-gen-buf-plugin-openapi

# Build directory
BUILD_DIR=./bin

# Go build flags
GO_BUILD_FLAGS=-v

# Installation directory (this should be in your PATH)
INSTALL_DIR=$(GOPATH)/bin

# Ensure build directory exists
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Build the binary
build: $(BUILD_DIR)
	go build $(GO_BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)

# Install the binary to GOPATH/bin for buf to find it
install: build-all
	mkdir -p $(INSTALL_DIR)
	cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/protoc-gen-buf-plugin-openapi

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)

# Test the plugin with buf
test: install
	cd .. && buf generate

PLUGIN_NAME := buf-plugin-openapi

generate:
	protoc --go_out=. --go_opt=paths=source_relative proto/annotations.proto

.PHONY: build-all
build-all:  build
