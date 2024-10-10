package reactive

import (
	"net/http"

	"ad.com/http/internal/logic/reactive"
	"ad.com/http/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := reactive.NewChatLogic(r.Context(), svcCtx)
		err := l.Chat()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
