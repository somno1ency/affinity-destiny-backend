package group

import (
	"net/http"

	"ad.com/http/internal/logic/group"
	"ad.com/http/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func GroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := group.NewGroupListLogic(r.Context(), svcCtx)
		resp, err := l.GroupList()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
