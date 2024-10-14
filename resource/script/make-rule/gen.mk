# Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

.PHONY: gen
gen: gen.api gen.model

.PHONY: gen.api
gen.api:
	@$(GOCTL) api format --dir ./http/api
	@$(GOCTL) api go -api ./http/api/v1/affinity_destiny.api --dir ./http --style go_zero
	@$(MAKE) cr.add

.PHONY: gen.model
gen.model:
	@$(GOCTL) model mysql ddl -s ./resource/ddl/user.sql -d ./pkg/repo/user -i '' --style go_zero
	@$(GOCTL) model mysql ddl -s ./resource/ddl/user_contact.sql -d ./pkg/repo/user_contact -i '' --style go_zero
	@$(GOCTL) model mysql ddl -s ./resource/ddl/group.sql -d ./pkg/repo/group -i '' --style go_zero
	@$(GOCTL) model mysql ddl -s ./resource/ddl/group_contact.sql -d ./pkg/repo/group_contact -i '' --style go_zero
	@$(GOCTL) model mysql ddl -s ./resource/ddl/category.sql -d ./pkg/repo/category -i '' --style go_zero
	@$(GOCTL) model mysql ddl -s ./resource/ddl/resource.sql -d ./pkg/repo/resource -i '' --style go_zero
	@$(MAKE) cr.add
