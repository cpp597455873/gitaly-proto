TARGET_DIR := $(CURDIR)/_build
TARGET_SETUP := $(TARGET_DIR)/.ok
BIN_BUILD_DIR := $(TARGET_DIR)/bin

unexport GOROOT
unexport GOBIN
export GO111MODULE = on

export PATH := $(TARGET_DIR)/bin:$(PATH)

# Developer Tools
PROTOC = $(TARGET_DIR)/protoc/bin/protoc
PROTOC_GEN_GO = $(BIN_BUILD_DIR)/protoc-gen-go
PROTOC_GEN_DOC = $(BIN_BUILD_DIR)/protoc-gen-doc
PROTOC_GEN_GITALY := $(BIN_BUILD_DIR)/protoc-gen-gitaly

.PHONY: all
all: generate

$(TARGET_SETUP):
	rm -rf $(TARGET_DIR)
	mkdir -p $(BIN_BUILD_DIR)
	touch $(TARGET_SETUP)

.PHONY: generate
generate: install-developer-tools go/gitalypb/*.pb.go
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
install-developer-tools: $(TARGET_SETUP) $(PROTOC) .ruby-bundle $(PROTOC_GEN_GITALY)

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

$(PROTOC_GEN_GO): $(TARGET_SETUP)
	cd go/internal && go build -o $@ github.com/golang/protobuf/protoc-gen-go

$(PROTOC_GEN_DOC): $(TARGET_SETUP)
	cd go/internal && go build -o $@ github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

$(PROTOC_GEN_GITALY): $(TARGET_SETUP)
	# Check if test protobuf stubs are stale
	cd go/internal && go build -o $@ gitlab.com/gitlab-org/gitaly-proto/go/internal/cmd/protoc-gen-gitaly

.PHONY: pb-go-stubs
pb-go-stubs: go/gitalypb/*.pb.go

go/gitalypb/%.pb.go: %.proto $(PROTOC) $(PROTOC_GEN_GO) $(PROTOC_GEN_GITALY)
	$(PROTOC) --gitaly_out=proto_dir=.,gitalypb_dir=go/gitalypb:. --go_out=paths=source_relative,plugins=grpc:./go/gitalypb -I$(shell pwd) *.proto

go/internal/linter/testdata/%.pb.go: go/internal/linter/testdata/%.proto $(PROTOC) $(PROTOC_GEN_GO) go/gitalypb/*.pb.go
	$(PROTOC) --go_out=paths=source_relative:. -I$(shell pwd) -I$(shell pwd)/go/internal/linter/testdata go/internal/linter/testdata/*.proto

.PHONY: test
test: test-go-pkg-opt
	cd go/internal && go test ./...

# test-go-pkg-opt checks if the go_package option is specified in all *.proto
# files
.PHONY: test-go-pkg-opt
test-go-pkg-opt:
	@for p in *.proto ; do \
		if ! grep -Fq "option go_package" $${p} ; then \
			echo "$${p} is missing the go_package option" ; \
			exit 1 ; \
		fi \
	done
