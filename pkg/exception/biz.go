package exception

import (
	"strconv"

	"ad.com/pkg/shared"
	"github.com/zeromicro/x/errors"
)

var (
	// 10_01_01 userService_userModule_errorType
	// 10_02_01 userService_groupModule_errorType
	// 10_03_01 userService_contactModule_errorType
	// 10_04_01 userService_chatModule_errorType
	// here Msg can be replaced by i18n to support multi-language
	UserNotFound               errors.CodeMsg = errors.CodeMsg{Code: 100101, Msg: "用户不存在..."}
	UserPasswordValidateFailed errors.CodeMsg = errors.CodeMsg{Code: 100101, Msg: "用户密码验证失败..."}
	UserPasswordUpdateFailed   errors.CodeMsg = errors.CodeMsg{Code: 100102, Msg: "用户密码更新失败..."}
	UserLoginTimeUpdateFailed  errors.CodeMsg = errors.CodeMsg{Code: 100103, Msg: "用户登录时间更新失败..."}
	UserUpdateFailed           errors.CodeMsg = errors.CodeMsg{Code: 100104, Msg: "用户信息更新失败..."}

	UserExists         errors.CodeMsg = errors.CodeMsg{Code: 100102, Msg: "用户已存在..."}
	UserNotMatch       errors.CodeMsg = errors.CodeMsg{Code: 100103, Msg: "用户名或密码错误..."}
	UserRegisterFailed errors.CodeMsg = errors.CodeMsg{Code: 100104, Msg: "用户注册失败..."}
	UserAddSelfFailed  errors.CodeMsg = errors.CodeMsg{Code: 100105, Msg: "不能添加自己为好友..."}
	UserAlreadyContact errors.CodeMsg = errors.CodeMsg{Code: 100106, Msg: "您已经是该用户的好友..."}
	UserAddFailed      errors.CodeMsg = errors.CodeMsg{Code: 100107, Msg: "好友添加失败..."}

	GroupNotFound       errors.CodeMsg = errors.CodeMsg{Code: 100201, Msg: "群不存在..."}
	GroupExists         errors.CodeMsg = errors.CodeMsg{Code: 100202, Msg: "群已存在..."}
	GroupCreateFailed   errors.CodeMsg = errors.CodeMsg{Code: 100203, Msg: "群创建失败..."}
	GroupQualityLimited errors.CodeMsg = errors.CodeMsg{Code: 100204, Msg: "每个人创建的群不能超过" + strconv.FormatInt(shared.GroupMaxQuantity, 10) + "个..."}
	GroupAlreadyIn      errors.CodeMsg = errors.CodeMsg{Code: 100205, Msg: "您已经是该群成员..."}
	GroupJoinFailed     errors.CodeMsg = errors.CodeMsg{Code: 100206, Msg: "加入群失败..."}
)
