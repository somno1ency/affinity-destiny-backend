# Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

GO := go
GOCTL := goctl
GOROOT := $(shell go env GOROOT)
ifeq ($(origin GOBIN), undefined)
	GOBIN := $(GOROOT)/bin
endif

.PHONY: go.update
go.updates: tool.verify.go-mod-outdated
	@$(GO) list -u -m -json all | go-mod-outdated -update -direct
