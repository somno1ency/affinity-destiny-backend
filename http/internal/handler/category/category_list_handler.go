package category

import (
	"net/http"

	"ad.com/http/internal/logic/category"
	"ad.com/http/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
