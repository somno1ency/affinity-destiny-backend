// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	contact "ad.com/http/internal/handler/contact"
	group "ad.com/http/internal/handler/group"
	reactive "ad.com/http/internal/handler/reactive"
	user "ad.com/http/internal/handler/user"
	"ad.com/http/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: contact.ContactAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/load/:id",
				Handler: contact.ContactLoadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/contacts"),
		rest.WithTimeout(10000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: group.GroupCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/join",
				Handler: group.GroupJoinHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/load/:id",
				Handler: group.GroupLoadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/groups"),
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
				Path:    "/:id",
				Handler: user.UserFindHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.UserRegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/users"),
		rest.WithTimeout(10000*time.Millisecond),
	)
}
