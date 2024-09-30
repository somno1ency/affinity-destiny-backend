package reactive

import (
	"net/http"

	"ad.com/resource/generated/go/internal/logic/reactive"
	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := reactive.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
