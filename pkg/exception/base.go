package exception

import "github.com/zeromicro/x/errors"

var (
	// 使用6位数字代表业务错误码,起始从100001开始
	// 10_00_01: 服务_模块_错误类型(模块00代表通用模块,所以base_error这个文件定义的都是00级别的错误)
	// 实际上base_err应该放到底层包中,所有的go项目引用即可,不用在每个项目中重复定义
	UnknownError errors.CodeMsg = errors.CodeMsg{Code: 100001}
)
