GITHUB_REPO := github.com/mi11km/monorepo-template
GO ?= go
GOOS := $(shell $(GO) env GOOS)
GOARCH := $(shell $(GO) env GOARCH)
BIN := $(ROOT)/bin/$(GOOS)_$(GOARCH)
GO_ENV ?= CGO_ENABLED=0 GOPRIVATE=$(GITHUB_REPO) GOBIN=$(BIN) TZ=Asia/Tokyo

# tools
$(shell mkdir -p $(BIN))

GOVULNCHECK_VERSION := 1.1.4
$(BIN)/govulncheck-$(GOVULNCHECK_VERSION):
	unlink $(BIN)/govulncheck || true
	$(GO_ENV) ${GO} install golang.org/x/vuln/cmd/govulncheck@v$(GOVULNCHECK_VERSION)
	mv $(BIN)/govulncheck $(BIN)/govulncheck-$(GOVULNCHECK_VERSION)
	ln -s $(BIN)/govulncheck-$(GOVULNCHECK_VERSION) $(BIN)/govulncheck

GRPCUI_VERSION := 1.4.3
$(BIN)/grpcui-$(GRPCUI_VERSION):
	unlink $(BIN)/grpcui || true
	$(GO_ENV) ${GO} install github.com/fullstorydev/grpcui/cmd/grpcui@v$(GRPCUI_VERSION)
	mv $(BIN)/grpcui $(BIN)/grpcui-$(GRPCUI_VERSION)
	ln -s $(BIN)/grpcui-$(GRPCUI_VERSION) $(BIN)/grpcui

.PHONY: install-tools
install-tools: $(BIN)/govulncheck-$(GOVULNCHECK_VERSION)
install-tools: $(BIN)/grpcui-$(GRPCUI_VERSION)
