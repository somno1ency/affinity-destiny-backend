package user_contact

import (
	"net/http"

	"ad.com/http/internal/logic/user_contact"
	"ad.com/http/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserContactListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user_contact.NewUserContactListLogic(r.Context(), svcCtx)
		resp, err := l.UserContactList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
