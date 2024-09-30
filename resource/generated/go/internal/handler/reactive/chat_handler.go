package reactive

import (
	"net/http"

	"ad.com/resource/generated/go/internal/logic/reactive"
	"ad.com/resource/generated/go/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := reactive.NewChatLogic(r.Context(), svcCtx)
		err := l.Chat()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
