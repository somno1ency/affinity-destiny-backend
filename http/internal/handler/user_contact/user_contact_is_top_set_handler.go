package user_contact

import (
	"net/http"

	"ad.com/http/internal/logic/user_contact"
	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserContactIsTopSetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserContactIsSet
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user_contact.NewUserContactIsTopSetLogic(r.Context(), svcCtx)
		err := l.UserContactIsTopSet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
