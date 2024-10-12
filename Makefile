# Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

include resource/script/make-rule/common.mk
include resource/script/make-rule/copyright.mk
include resource/script/make-rule/dependency.mk
include resource/script/make-rule/gen.mk
include resource/script/make-rule/go.mk 
include resource/script/make-rule/tool.mk

# build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: dependency update add-cr

.PHONY: dependency
dependency:
	@$(MAKE) dependency.run

.PHONY: update
update:
	@$(MAKE) go.update

.PHONY: tidy
tidy:
	@$(MAKE) dependency.package

.PHONY: gen
gen:
	@$(MAKE) gen.api
	@$(MAKE) gen.model

.PHONY: gen-api
gen-api:
	@$(MAKE) gen.api

.PHONY: gen-model
gen-model:
	@$(MAKE) gen.model

.PHONY: add-cr
add-cr:
	@$(MAKE) cr.add
