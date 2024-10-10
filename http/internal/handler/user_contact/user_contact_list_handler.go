package user_contact

import (
	"net/http"

	"ad.com/http/internal/logic/user_contact"
	"ad.com/http/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func UserContactListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user_contact.NewUserContactListLogic(r.Context(), svcCtx)
		resp, err := l.UserContactList()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
