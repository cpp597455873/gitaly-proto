TARGET_DIR := $(CURDIR)/_build
TARGET_SETUP := $(TARGET_DIR)/.ok
BIN_BUILD_DIR := $(TARGET_DIR)/bin

export GOPATH := $(TARGET_DIR)
export PATH := $(GOPATH)/bin:$(PATH)

# Developer Tools
PROTOC = $(TARGET_DIR)/protoc/bin/protoc
PROTOC_GEN_RUBY = $(BIN_BUILD_DIR)/grpc_tools_ruby_protoc
PROTOC_GEN_GO = $(BIN_BUILD_DIR)/protoc-gen-go

.PHONY: all
all: generate

$(TARGET_SETUP):
	rm -rf $(TARGET_DIR)
	mkdir -p $(BIN_BUILD_DIR)
	touch $(TARGET_SETUP)

.PHONY: build
generate: install-developer-tools
	_support/generate-from-proto

.PHONY: clean
clean:
	rm -rf $(TARGET_DIR)

.PHONY: release
release: install-developer-tools
ifeq ($(version), "")
	$error("Please run 'make release version=x.y.z'")
endif
	_support/release $(version)

.PHONY: check-grpc-proto-clients
check-grpc-proto-clients: install-developer-tools
	_support/check-grpc-proto-clients

.PHONY: install-developer-tools
install-developer-tools: $(TARGET_SETUP) $(PROTOC) $(PROTOC_GEN_GO) $(PROTOC_GEN_RUBY)

$(PROTOC): $(TARGET_SETUP)
	_support/install-protoc

$(PROTOC_GEN_GO): $(TARGET_SETUP)
	go get -v github.com/golang/protobuf/protoc-gen-go

$(PROTOC_GEN_RUBY): $(TARGET_SETUP) _support/Gemfile
	bundle install --gemfile=_support/Gemfile --binstubs=$(BIN_BUILD_DIR)
