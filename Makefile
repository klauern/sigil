MAIN_PKG := github.com/danmx/sigil

IS_LINUX := $(if $(filter-out Darwin,$(shell uname -s)),true,)

GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

GOLANGCI_LINT_BIN := $(GOBIN)/golangci-lint
GOLANGCI_LINT_VERSION := 1.32.2

GOTESTSUM_BIN := $(GOBIN)/gotestsum
GOTESTSUM_VERSION := 0.5.4

export GOFLAGS := -v

.PHONY: build clean debug debug_remote enable_arturos ensure_deps get lint lint_fix run_dev test test_ci

default: build

build:
	go build $(MAIN_PKG)

ensure_deps:
	go mod tidy
	go mod vendor

lint: $(GOLANGCI_LINT_BIN)
	"$(GOLANGCI_LINT_BIN)" run

lint_fix: $(GOLANGCI_LINT_BIN)
	"$(GOLANGCI_LINT_BIN)" run --fix

generate:
	go generate ./...

test: $(GOTESTSUM_BIN)
	@CGO_ENABLED=1 "$(GOTESTSUM_BIN)" -- \
		-race \
			./...

test_ci: build lint test

clean:
	go clean $(MAIN_PKG) && rm -f ./sigil

$(GOLANGCI_LINT_BIN):
	@curl --fail --location --silent https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b "$(GOBIN)" "v$(GOLANGCI_LINT_VERSION)"

$(GOTESTSUM_BIN):
	$(eval path := v$(GOTESTSUM_VERSION)/gotestsum_$(GOTESTSUM_VERSION)_$(if $(IS_LINUX),linux,darwin)_amd64.tar.gz)
	@curl --fail --location --silent "https://github.com/gotestyourself/gotestsum/releases/download/$(path)" | \
		tar -xz -C "$(GOBIN)" -- gotestsum
