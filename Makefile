.DEFAULT_GOAL := proto-lint

BIN := .bin
$(BIN):
	@mkdir -p $@

# https://docs.buf.build/
# changing BUF_VERSION will automatically download and use the specified version.
BUF_VERSION = 0.36.0
OS = $(shell uname -s)
ARCH = $(shell uname -m)
BUF_URL = https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(OS)-$(ARCH)
# use BUF_VERSION_BIN as a bin prerequisite, not "buf", so the correct version will be used.
BUF_VERSION_BIN = buf-$(BUF_VERSION)
$(BIN)/$(BUF_VERSION_BIN): | $(BIN)
	@echo "downloading buf $(BUF_VERSION)"
	@curl -sSL $(BUF_URL) -o $@
	@chmod +x $@

PROTO_ROOT := proto
PROTO_FILES = $(shell find ./$(PROTO_ROOT) -name "*.proto")
proto-lint: $(PROTO_FILES) $(BIN)/$(BUF_VERSION_BIN)
	@$(BIN)/$(BUF_VERSION_BIN) lint
