# build all by default, even if it's not first
.DEFAULT_GOAL := all

# ==============================================================================
# Includes
include resource/script/make-rule/go.mk 

.PHONY: all
all: tidy gen cr format lint cover build

.PHONY: tidy
tidy:
	@$(GO) mod tidy

.PHONY: gen.api
gen.api:
	@$(GO_CTL) api go -api ./http/api/v1/affinity_destiny.api --dir ./http --style go_zero

.PHONY: gen.model
gen.model:
	@$(GO_CTL) model mysql ddl -s ./resource/ddl/user.sql -d ./pkg/repo/user -i '' --style go_zero
	@$(GO_CTL) model mysql ddl -s ./resource/ddl/user_contact.sql -d ./pkg/repo/user_contact -i '' --style go_zero
	@$(GO_CTL) model mysql ddl -s ./resource/ddl/group.sql -d ./pkg/repo/group -i '' --style go_zero
	@$(GO_CTL) model mysql ddl -s ./resource/ddl/group_contact.sql -d ./pkg/repo/group_contact -i '' --style go_zero
	@$(GO_CTL) model mysql ddl -s ./resource/ddl/category.sql -d ./pkg/repo/category -i '' --style go_zero
	@$(GO_CTL) model mysql ddl -s ./resource/ddl/resource.sql -d ./pkg/repo/resource -i '' --style go_zero

.PHONY: format.api
format.api:
	@$(GO_CTL) api format --dir ./http/api