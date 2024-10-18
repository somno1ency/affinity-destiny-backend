// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package exception

import "github.com/zeromicro/x/errors"

var (
	// use 6-digit number represents error code, start from 100001
	// 100_00_001: service_module_errorType(module 00 represents common module, so this file defines the 00 level error)
	// actually base error need to put to infra package, and referenced by all go project, not duplicated defines in each project
	UnknownError         errors.CodeMsg = errors.CodeMsg{Code: 100_00_001}
	CodeSendFailed       errors.CodeMsg = errors.CodeMsg{Code: 100_00_002, Msg: "验证码发送失败..."}
	CodeValidateFailed   errors.CodeMsg = errors.CodeMsg{Code: 100_00_003, Msg: "验证码校验失败..."}
	UnknownAscFieldError errors.CodeMsg = errors.CodeMsg{Code: 100_00_004, Msg: "排序字段不存在..."}
	NoPermissionError    errors.CodeMsg = errors.CodeMsg{Code: 100_00_005, Msg: "没有权限执行..."}
)
