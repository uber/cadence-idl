.DEFAULT_GOAL := all
.NOTPARALLEL: # not parallel friendly for simplicity

# CAUTION:
# you are NOT expected to run individual targets by hand.
# they will not re-build correctly.
# for normal changes, just use `make` and `make clean` when necessary.

# ###############################
#         general helpers
# ###############################

# M1 macs may need to switch back to x86, until arm releases are available
EMULATE_X86 =
ifeq ($(shell uname -sm),Darwin arm64)
EMULATE_X86 = arch -x86_64
endif

OS = $(shell uname -s)
ARCH = $(shell $(EMULATE_X86) uname -m)

# current (when committed) version of Go used in CI, and ideally also our docker images.
# this generally does not matter, but can impact goimports output.
# for maximum stability, make sure you use the same version as CI uses.
#
# this can _likely_ remain a major version, as fmt output does not tend to change in minor versions,
# which will allow findstring to match any minor version.
EXPECTED_GO_VERSION := go1.17
CURRENT_GO_VERSION := $(shell go version)
ifeq (,$(findstring $(EXPECTED_GO_VERSION),$(CURRENT_GO_VERSION)))
# if you are seeing this warning: consider using https://github.com/travis-ci/gimme to pin your version
$(warning Caution: you are not using CI's go version. Expected: $(EXPECTED_GO_VERSION), current: $(CURRENT_GO_VERSION))
$(warning If you have recently changed your version: make sure to clean and rebuild, as pre-built tools generally inherit Go's formatting for that version.)
endif

# set a V=1 env var for verbose output. V=0 (or unset) disables.
# this is used to make two verbose flags:
# - $Q, to replace ALL @ use, so CI can be reliably verbose
# - $(verbose), to forward verbosity flags to commands via `$(if $(verbose),-v)` or similar
#
# SHELL='bash -x' is useful too, but can be more confusing to understand.
V ?= 0
ifneq (0,$(V))
verbose := 1
Q :=
else
verbose :=
Q := @
endif

# and enforce ^ that rule: grep the makefile for line-starting @ use, error if any exist.
# limit to one match because multiple look too weird.
_BAD_AT_USE=$(shell grep -n -m1 '^\s*@' $(MAKEFILE_LIST))
ifneq (,$(_BAD_AT_USE))
$(warning Makefile cannot use @ to silence commands, use $$Q instead:)
$(warning found on line $(_BAD_AT_USE))
$(error fix that line and try again)
endif

BIN := .bin
$(BIN):
	$Q mkdir -p $@

# helper for executing bins that need other bins, just `$(BIN_PATH) the_command ...`
# I'd recommend not exporting this in general, to reduce the chance of accidentally using non-versioned tools.
BIN_PATH := PATH="$(abspath $(BIN)):$$PATH"

# builds a go-gettable tool, versioned by internal/tools/go.mod, and installs it into
# the build folder, named the same as the last portion of the URL or the second arg.
define go_build_tool
$Q echo "building $(or $(2), $(notdir $(1))) from internal/tools/go.mod..."
$Q go build -mod=readonly -modfile=internal/tools/go.mod -o $(BIN)/$(or $(2), $(notdir $(1))) $(1)
endef

# prefix lines in license with // comments, and prepend it to all files in the "go" generated dir
license:
	$Q echo "prefixing license"
	$Q export TMP="$$(mktemp)"; \
		sed 's/^/\/\/ /' LICENSE > "$$TMP"; \
		echo >> $$TMP; \
		find go -type f -exec sh -c 'cat $$TMP $$1 > $$1.tmp; mv $$1.tmp $$1' sh {} \;

# caution: matches the filename on case-insensitive systems
.PHONY: license

# ###############################
#         protoc stuff
# ###############################

# https://docs.buf.build/
# changing BUF_VERSION will automatically download and use the specified version.
BUF_VERSION = 0.36.0
BUF_URL = https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(OS)-$(ARCH)
# use BUF_VERSION_BIN as a bin prerequisite, not "buf", so the correct version will be used.
BUF_VERSION_BIN = buf-$(BUF_VERSION)
$(BIN)/$(BUF_VERSION_BIN): | $(BIN)
	$Q echo "downloading buf $(BUF_VERSION)"
	$Q curl -sSL $(BUF_URL) -o $@
	$Q chmod +x $@

PROTO_ROOT := proto
PROTO_FILES = $(shell find ./$(PROTO_ROOT) -name "*.proto")
PROTO_DIRS = $(sort $(dir $(PROTO_FILES)))
proto-lint: $(PROTO_FILES) $(BIN)/$(BUF_VERSION_BIN)
	$Q $(BIN)/$(BUF_VERSION_BIN) lint

# https://www.grpc.io/docs/languages/go/quickstart/
# protoc-gen-gogofast (yarpc) are versioned via tools.go + go.mod (built above) and will be rebuilt as needed.
# changing PROTOC_VERSION will automatically download and use the specified version
PROTOC_VERSION = 3.14.0
PROTOC_URL = https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(subst Darwin,osx,$(OS))-$(ARCH).zip
# the zip contains an /include folder that we need to use to learn the well-known types
PROTOC_UNZIP_DIR = $(BIN)/protoc-$(PROTOC_VERSION)-zip
# use PROTOC_VERSION_BIN as a bin prerequisite, not "protoc", so the correct version will be used.
# otherwise this must be a .PHONY rule, or the buf bin / symlink could become out of date.
PROTOC_VERSION_BIN = protoc-$(PROTOC_VERSION)
$(BIN)/$(PROTOC_VERSION_BIN): | $(BIN)
	$Q echo "downloading protoc $(PROTOC_VERSION)"
	$Q # recover from partial success
	$Q rm -rf $(BIN)/protoc.zip $(PROTOC_UNZIP_DIR)
	$Q # download, unzip, copy to a normal location
	$Q curl -sSL $(PROTOC_URL) -o $(BIN)/protoc.zip
	$Q unzip -q $(BIN)/protoc.zip -d $(PROTOC_UNZIP_DIR)
	$Q cp $(PROTOC_UNZIP_DIR)/bin/protoc $@

$(BIN)/protoc-gen-gogofast: internal/tools/go.mod $(BIN)/go_mod_check
	$(call go_build_tool,github.com/gogo/protobuf/protoc-gen-gogofast)

$(BIN)/protoc-gen-yarpc-go: internal/tools/go.mod $(BIN)/go_mod_check
	$(call go_build_tool,go.uber.org/yarpc/encoding/protobuf/protoc-gen-yarpc-go)

PROTO_GO_OUT := go/proto
LICENSE_GO := LICENSE.go

define build_proto
$Q echo "protoc for $(1)"
$Q $(EMULATE_X86) $(BIN)/$(PROTOC_VERSION_BIN) \
	--plugin $(BIN)/protoc-gen-gogofast \
	--plugin $(BIN)/protoc-gen-yarpc-go \
	-I=$(PROTO_ROOT) \
	-I=$(PROTOC_UNZIP_DIR)/include \
	--gogofast_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,paths=source_relative:$(PROTO_GO_OUT) \
	--yarpc-go_out=$(PROTO_GO_OUT) \
	$$(find $(1) -name '*.proto')
endef

proto: $(PROTO_FILES) $(BIN)/$(PROTOC_VERSION_BIN) $(BIN)/protoc-gen-gogofast $(BIN)/protoc-gen-yarpc-go
	$Q # clean up previous codegen
	$Q rm -rf $(PROTO_GO_OUT)
	$Q mkdir -p $(PROTO_GO_OUT)
	$Q $(call build_proto,proto/uber/cadence/admin/v1)
	$Q $(call build_proto,proto/uber/cadence/api/v1)
	$Q # move go/proto/uber/cadence/... to go/proto/... to match the import path we want, and clean up the empty dir
	$Q mv $(PROTO_GO_OUT)/uber/cadence/* $(PROTO_GO_OUT)
	$Q rm -rf $(PROTO_GO_OUT)/uber

# ###############################
#         thrift stuff
# ###############################

$(BIN)/thriftrw: internal/tools/go.mod $(BIN)/go_mod_check
	$(call go_build_tool,go.uber.org/thriftrw)

$(BIN)/thriftrw-plugin-yarpc: internal/tools/go.mod $(BIN)/go_mod_check
	$(call go_build_tool,go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc)

SYNCED_MODULES = go.uber.org/thriftrw go.uber.org/yarpc github.com/gogo/protobuf
# ensures mod files are in sync for critical packages.
# this stores a bookkeeping file in $(BIN) for simplicity, so it does not force tool-rebuilding every time.
$(BIN)/go_mod_check: go.mod internal/tools/go.mod
	$Q # ensure both mods have the same core library versions
	$Q export TOP="$$(mktemp -t top)"; \
		export TOOL="$$(mktemp -t tool)"; \
		2>&1 go list -m -mod=readonly $(SYNCED_MODULES) | sort > "$$TOP"; \
		2>&1 go list -m -mod=readonly -modfile=internal/tools/go.mod $(SYNCED_MODULES) | sort > "$$TOOL"; \
		diff "$$TOP" "$$TOOL" || >&2 echo "top-level and tool dependencies out of sync"
	$Q touch $@

define build_thrift
$Q echo 'thriftrw for $(1)'
$Q $(BIN_PATH) $(BIN)/thriftrw \
	--plugin=yarpc \
	--pkg-prefix=github.com/uber/cadence-idl/go/thrift \
	--out=go/thrift \
	$(1)
endef

# thrifts are generated transitively, so only the roots of trees need to be built.
thrift: $(BIN)/thriftrw $(BIN)/thriftrw-plugin-yarpc
	$(call build_thrift,thrift/cadence.thrift)
	$(call build_thrift,thrift/shadower.thrift)

.PHONY: thrift

# ###############################
#         primary targets
# ###############################

all: $(BIN) proto-lint proto thrift license

# generally not necessary unless we change library versions, but this DOES impact codegen
# formatting because it inherits that from the version of Go used to build tools.
clean:
	rm -rf $(BIN)
