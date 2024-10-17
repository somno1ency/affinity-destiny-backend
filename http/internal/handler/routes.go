// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	admin_resource "ad.com/http/internal/handler/admin_resource"
	category "ad.com/http/internal/handler/category"
	group "ad.com/http/internal/handler/group"
	group_contact "ad.com/http/internal/handler/group_contact"
	reactive "ad.com/http/internal/handler/reactive"
	resource "ad.com/http/internal/handler/resource"
	user "ad.com/http/internal/handler/user"
	user_contact "ad.com/http/internal/handler/user_contact"
	"ad.com/http/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: admin_resource.ResourceListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: admin_resource.ResourceCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: admin_resource.ResourceUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: admin_resource.ResourceDeleteHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/admin/resources"),
		rest.WithTimeout(10000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: category.CategoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: category.CategoryCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: category.CategoryUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: category.CategoryDeleteHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/categories"),
		rest.WithTimeout(10000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: group.GroupListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: group.GroupCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: group.GroupUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/disband",
				Handler: group.GroupDisbandHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/transfer",
				Handler: group.GroupTransferHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/groups"),
		rest.WithTimeout(10000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/:groupId",
				Handler: group_contact.GroupContactListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:groupId",
				Handler: group_contact.GroupContactAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:groupId/join",
				Handler: group_contact.GroupContactJoinHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:groupId/leave",
				Handler: group_contact.GroupContactLeaveHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: group_contact.GroupContactDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/background",
				Handler: group_contact.GroupContactBackgroundSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/category",
				Handler: group_contact.GroupContactCategorySetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/is-disturb",
				Handler: group_contact.GroupContactIsDisturbSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/is-show-nickname",
				Handler: group_contact.GroupContactIsShowNicknameSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/is-top",
				Handler: group_contact.GroupContactIsTopSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/nickname",
				Handler: group_contact.GroupContactNicknameSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/remark",
				Handler: group_contact.GroupContactRemarkSetHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/group-contacts"),
		rest.WithTimeout(10000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/chat",
				Handler: reactive.ChatHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/upload",
				Handler: reactive.UploadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
		rest.WithTimeout(10000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: resource.ResourceListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/resources"),
		rest.WithTimeout(10000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: user.UserFindHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: user.UserUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id/password",
				Handler: user.PasswordIsSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/password",
				Handler: user.PasswordSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/send-code",
				Handler: user.CodeSendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/validate-code",
				Handler: user.CodeValidateHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/users"),
		rest.WithTimeout(10000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: user_contact.UserContactListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id",
				Handler: user_contact.UserContactAddHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: user_contact.UserContactDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/background",
				Handler: user_contact.UserContactBackgroundSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/category",
				Handler: user_contact.UserContactCategorySetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/is-disturb",
				Handler: user_contact.UserContactIsDisturbSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/is-remind",
				Handler: user_contact.UserContactIsRemindSetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id/is-top",
				Handler: user_contact.UserContactIsTopSetHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user-contacts"),
		rest.WithTimeout(10000*time.Millisecond),
	)
}
