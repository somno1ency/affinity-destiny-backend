package group_contact

import (
	"net/http"

	"ad.com/http/internal/logic/group_contact"
	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func GroupContactIsShowNicknameSetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupContactIsSet
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := group_contact.NewGroupContactIsShowNicknameSetLogic(r.Context(), svcCtx)
		err := l.GroupContactIsShowNicknameSet(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
