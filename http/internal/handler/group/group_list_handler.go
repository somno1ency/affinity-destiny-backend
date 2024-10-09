package group

import (
	"net/http"

	"ad.com/http/internal/logic/group"
	"ad.com/http/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := group.NewGroupListLogic(r.Context(), svcCtx)
		resp, err := l.GroupList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
