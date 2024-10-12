# Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

.PHONY: cr.add
cr.add:
	@addlicense -v -f $(RESOURCE_DIR)/doc/copy-right.txt $(ROOT_DIR) --skip-dirs=.history,.vscode,githooks --skip-files=.api,.txt