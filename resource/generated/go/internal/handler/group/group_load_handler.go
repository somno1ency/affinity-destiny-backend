package group

import (
	"net/http"

	"ad.com/resource/generated/go/internal/logic/group"
	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupLoadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGroupLoadLogic(r.Context(), svcCtx)
		resp, err := l.GroupLoad(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
