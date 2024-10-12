# Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

.PHONY: dependency.run
dependency.run: dependency.package dependency.tool

.PHONY: dependency.package
dependency.package:
	@$(GO) mod tidy

.PHONY: dependency.tool
dependency.tool: dependency.tool.blocker dependency.tool.critical

.PHONY: dependency.tool.blocker
dependency.tool.blocker: $(addprefix tool.verify., $(BLOCKER_TOOL))

.PHONY: dependency.tool.critical
dependency.tool.critical: $(addprefix tool.verify., $(CRITICAL_TOOL))
