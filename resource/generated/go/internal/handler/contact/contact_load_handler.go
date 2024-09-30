package contact

import (
	"net/http"

	"ad.com/resource/generated/go/internal/logic/contact"
	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ContactLoadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := contact.NewContactLoadLogic(r.Context(), svcCtx)
		resp, err := l.ContactLoad(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
