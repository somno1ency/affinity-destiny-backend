package group

import (
	"net/http"

	"ad.com/resource/generated/go/internal/logic/group"
	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupJoinHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupJoinReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGroupJoinLogic(r.Context(), svcCtx)
		err := l.GroupJoin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
