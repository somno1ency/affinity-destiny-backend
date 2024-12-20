// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package exception

import (
	"github.com/zeromicro/x/errors"
)

var (
	// 101_01_01 user_user_error
	// 101_02_01 user_group_error
	// 101_03_01 user_category_error
	// 101_04_01 user_resource_error
	// here Msg can be replaced by i18n to support multi-language
	UserNotFound               errors.CodeMsg = errors.CodeMsg{Code: 101_01_001, Msg: "用户不存在..."}
	UserPasswordValidateFailed errors.CodeMsg = errors.CodeMsg{Code: 101_01_002, Msg: "用户密码验证失败..."}
	UserPasswordUpdateFailed   errors.CodeMsg = errors.CodeMsg{Code: 101_01_003, Msg: "用户密码更新失败..."}
	UserLoginTimeUpdateFailed  errors.CodeMsg = errors.CodeMsg{Code: 101_01_004, Msg: "用户登录时间更新失败..."}
	UserUpdateFailed           errors.CodeMsg = errors.CodeMsg{Code: 101_01_005, Msg: "用户信息更新失败..."}
	UserRegisterFailed         errors.CodeMsg = errors.CodeMsg{Code: 101_01_006, Msg: "用户注册失败..."}
	UserContactListQueryFailed errors.CodeMsg = errors.CodeMsg{Code: 101_01_007, Msg: "好友列表查询失败..."}
	UserContactCountFailed     errors.CodeMsg = errors.CodeMsg{Code: 101_01_008, Msg: "好友数量查询失败..."}
	UserContactAlreadyExist    errors.CodeMsg = errors.CodeMsg{Code: 101_01_009, Msg: "该用户已经是您的好友..."}
	UserContactAddSelfFailed   errors.CodeMsg = errors.CodeMsg{Code: 101_01_010, Msg: "不能添加自己为好友..."}
	UserContactCreateFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_01_011, Msg: "好友添加失败..."}
	UserContactUpdateFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_01_012, Msg: "好友信息更新失败..."}
	UserContactDeleteFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_01_013, Msg: "好友删除失败..."}
	UserContactNotTargetFriend errors.CodeMsg = errors.CodeMsg{Code: 101_01_014, Msg: "您不是对方的好友..."}
	UserContactTargetNotFriend errors.CodeMsg = errors.CodeMsg{Code: 101_01_015, Msg: "对方不是您的好友..."}
	UserContactTargetInvite    errors.CodeMsg = errors.CodeMsg{Code: 101_01_016, Msg: "对方已经申请成为您的好友,待您处理..."}
	UserContactSelfInvite      errors.CodeMsg = errors.CodeMsg{Code: 101_01_017, Msg: "您已经申请成为对方好友,待对方处理..."}

	GroupCreateFailed           errors.CodeMsg = errors.CodeMsg{Code: 101_02_001, Msg: "群创建失败..."}
	GroupNotFound               errors.CodeMsg = errors.CodeMsg{Code: 101_02_002, Msg: "群不存在..."}
	GroupUpdateFailed           errors.CodeMsg = errors.CodeMsg{Code: 101_02_003, Msg: "群更新失败..."}
	GroupListQueryFailed        errors.CodeMsg = errors.CodeMsg{Code: 101_02_004, Msg: "群列表查询失败..."}
	GroupCountFailed            errors.CodeMsg = errors.CodeMsg{Code: 101_02_005, Msg: "群数量查询失败..."}
	GroupDisbandFailed          errors.CodeMsg = errors.CodeMsg{Code: 101_02_006, Msg: "群解散失败..."}
	GroupTransferFailed         errors.CodeMsg = errors.CodeMsg{Code: 101_02_007, Msg: "群转让失败..."}
	GroupTransferToSelfError    errors.CodeMsg = errors.CodeMsg{Code: 101_02_008, Msg: "群转让失败,不能转让给自己..."}
	GroupContactUserIdsEmpty    errors.CodeMsg = errors.CodeMsg{Code: 101_02_009, Msg: "好友用户不能为空..."}
	GroupContactAllUserNotExist errors.CodeMsg = errors.CodeMsg{Code: 101_02_010, Msg: "用户均不存在..."}
	GroupContactsCreateFailed   errors.CodeMsg = errors.CodeMsg{Code: 101_02_011, Msg: "群成员添加失败..."}
	GroupContactNotFound        errors.CodeMsg = errors.CodeMsg{Code: 101_02_012, Msg: "群成员不存在..."}
	GroupContactUpdateFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_02_013, Msg: "群成员信息更新失败..."}
	GroupContactListQueryFailed errors.CodeMsg = errors.CodeMsg{Code: 101_02_014, Msg: "群成员列表查询失败..."}
	GroupContactDeleteFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_02_015, Msg: "群成员删除失败..."}
	GroupContactAlreadyExists   errors.CodeMsg = errors.CodeMsg{Code: 101_02_016, Msg: "群成员已存在..."}
	GroupContactJoinFailed      errors.CodeMsg = errors.CodeMsg{Code: 101_02_017, Msg: "加入群失败..."}

	CategoryCreateFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_03_001, Msg: "分组创建失败..."}
	CategoryDeleteFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_03_002, Msg: "分组删除失败..."}
	CategoryNotFound        errors.CodeMsg = errors.CodeMsg{Code: 101_03_003, Msg: "分组不存在..."}
	CategoryUpdateFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_03_004, Msg: "分组更新失败..."}
	CategoryListQueryFailed errors.CodeMsg = errors.CodeMsg{Code: 101_03_005, Msg: "分组列表查询失败..."}
	CategoryCountFailed     errors.CodeMsg = errors.CodeMsg{Code: 101_03_006, Msg: "分组数量查询失败..."}

	ResourceCreateFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_04_001, Msg: "资源创建失败..."}
	ResourceDeleteFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_04_002, Msg: "资源删除失败..."}
	ResourceNotFound        errors.CodeMsg = errors.CodeMsg{Code: 101_04_003, Msg: "资源不存在..."}
	ResourceUpdateFailed    errors.CodeMsg = errors.CodeMsg{Code: 101_04_004, Msg: "资源更新失败..."}
	ResourceListQueryFailed errors.CodeMsg = errors.CodeMsg{Code: 101_04_005, Msg: "资源列表查询失败..."}
	ResourceCountFailed     errors.CodeMsg = errors.CodeMsg{Code: 101_04_006, Msg: "资源数量查询失败..."}
)
