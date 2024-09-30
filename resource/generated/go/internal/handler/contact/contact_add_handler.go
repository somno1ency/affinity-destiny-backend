package contact

import (
	"net/http"

	"ad.com/resource/generated/go/internal/logic/contact"
	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ContactAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ContactAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := contact.NewContactAddLogic(r.Context(), svcCtx)
		err := l.ContactAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
