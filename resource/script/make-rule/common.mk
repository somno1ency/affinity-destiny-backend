# Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

SHELL := /bin/bash

MAKE_RULE_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# init dir
ifeq ($(origin RESOURCE_DIR), undefined)
RESOURCE_DIR := $(abspath $(shell cd $(MAKE_RULE_DIR)/../.. && pwd -P))
endif
ifeq ($(origin ROOT_DIR), undefined)
ROOT_DIR := $(abspath $(shell cd $(RESOURCE_DIR)/.. && pwd -P))
endif

COPY_GITHOOK:=$(shell cp -f $(RESOURCE_DIR)/githooks/* .git/hooks/)

# specify tools severity, include: BLOCKER_TOOL, CRITICAL_TOOL, TRIVIAL_TOOL
# missing BLOCKER_TOOL can cause the CI flow execution failed, i.e
# Missing CRITICAL_TOOL can lead to some necessary operations faile
BLOCKER_TOOL ?= golines go-junit-report addlicense goimports
CRITICAL_TOOL ?= mockgen gotests go-mod-outdated protoc-gen-go go-gitlint
