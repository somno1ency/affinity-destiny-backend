package exception

import (
	"strconv"

	global "ad.com/pkg/const"
	"github.com/zeromicro/x/errors"
)

var (
	// 10_01_01 user服务_user模块_错误序号
	// 10_02_01 user服务_group模块_错误序号
	// 10_03_01 user服务_contact模块_错误序号
	// 10_04_01 user服务_chat模块_错误序号
	// 这里Msg可根据环境中的语言参数读取对应的语言版本从而实现多语言化
	UserNotFound       errors.CodeMsg = errors.CodeMsg{Code: 100101, Msg: "用户不存在..."}
	UserExists         errors.CodeMsg = errors.CodeMsg{Code: 100102, Msg: "用户已存在..."}
	UserNotMatch       errors.CodeMsg = errors.CodeMsg{Code: 100103, Msg: "用户名或密码错误..."}
	UserRegisterFailed errors.CodeMsg = errors.CodeMsg{Code: 100104, Msg: "用户注册失败..."}
	UserAddSelfFailed  errors.CodeMsg = errors.CodeMsg{Code: 100105, Msg: "不能添加自己为好友..."}
	UserAlreadyContact errors.CodeMsg = errors.CodeMsg{Code: 100106, Msg: "您已经是该用户的好友..."}
	UserAddFailed      errors.CodeMsg = errors.CodeMsg{Code: 100107, Msg: "好友添加失败..."}

	GroupNotFound       errors.CodeMsg = errors.CodeMsg{Code: 100201, Msg: "群不存在..."}
	GroupExists         errors.CodeMsg = errors.CodeMsg{Code: 100202, Msg: "群已存在..."}
	GroupCreateFailed   errors.CodeMsg = errors.CodeMsg{Code: 100203, Msg: "群创建失败..."}
	GroupQualityLimited errors.CodeMsg = errors.CodeMsg{Code: 100204, Msg: "每个人创建的群不能超过" + strconv.FormatInt(global.GroupMaxQuantity, 10) + "个..."}
	GroupAlreadyIn      errors.CodeMsg = errors.CodeMsg{Code: 100205, Msg: "您已经是该群成员..."}
	GroupJoinFailed     errors.CodeMsg = errors.CodeMsg{Code: 100206, Msg: "加入群失败..."}
)
