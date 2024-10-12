# Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

TOOL ?=$(BLOCKER_TOOL) $(CRITICAL_TOOL) $(TRIVIAL_TOOL)

.PHONY: tool.install
tool.install: $(addprefix tool.install., $(TOOL))

.PHONY: tool.install.%
tool.install.%:
	@echo "===========> Installing $*"
	@$(MAKE) install.$*

.PHONY: tool.verify.%
tool.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tool.install.$*; fi

.PHONY: install.golines
install.golines:
	@$(GO) install github.com/segmentio/golines@latest

.PHONY: install.go-junit-report
install.go-junit-report:
	@$(GO) install github.com/jstemmer/go-junit-report@latest

.PHONY: install.addlicense
install.addlicense:
	@$(GO) install github.com/marmotedu/addlicense@latest

.PHONY: install.goimports
install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: install.mockgen
install.mockgen:
	@$(GO) install github.com/golang/mock/mockgen@latest

.PHONY: install.gotests
install.gotests:
	@$(GO) install github.com/cweill/gotests/gotests@latest

.PHONY: install.go-mod-outdated
install.go-mod-outdated:
	@$(GO) install github.com/psampaz/go-mod-outdated@latest

.PHONY: install.protoc-gen-go
install.protoc-gen-go:
	@$(GO) install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: install.go-gitlint
install.go-gitlint:
	@$(GO) install github.com/marmotedu/go-gitlint/cmd/go-gitlint@latest
