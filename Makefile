TARGET_DIR := $(CURDIR)/_build
TARGET_SETUP := $(TARGET_DIR)/.ok
BIN_BUILD_DIR := $(TARGET_DIR)/bin

unexport GOROOT
unexport GOBIN

export GOPATH := $(TARGET_DIR)
export PATH := $(GOPATH)/bin:$(PATH)

# Developer Tools
PROTOC = $(TARGET_DIR)/protoc/bin/protoc
PROTOC_GEN_GO = $(BIN_BUILD_DIR)/protoc-gen-go
PROTOC_GEN_DOC = $(BIN_BUILD_DIR)/protoc-gen-doc

.PHONY: all
all: generate

$(TARGET_SETUP):
	rm -rf $(TARGET_DIR)
	mkdir -p $(BIN_BUILD_DIR)
	touch $(TARGET_SETUP)

.PHONY: generate
generate: install-developer-tools
	_support/generate-from-proto

.PHONY: clean
clean:
	rm -rf $(TARGET_DIR) public .ruby-bundle

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
install-developer-tools: $(TARGET_SETUP) $(PROTOC) .ruby-bundle

.PHONY: docs
docs: $(TARGET_SETUP) $(PROTOC_GEN_DOC) $(PROTOC)
	rm -rf public && mkdir public
	$(PROTOC) --doc_out=./public \
	          --plugin=protoc-gen-doc=$(PROTOC_GEN_DOC) \
	          *.proto

$(PROTOC): $(TARGET_SETUP)
	_support/install-protoc

.ruby-bundle: $(TARGET_SETUP) _support/Gemfile.lock
	bundle install --gemfile=_support/Gemfile --binstubs=$(BIN_BUILD_DIR)
	touch $@	

$(PROTOC_GEN_DOC): $(TARGET_SETUP)
	go get -v -u github.com/pseudomuto/protoc-gen-doc/cmd/...
