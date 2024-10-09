package category

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryDeleteLogic {
	return &CategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryDeleteLogic) CategoryDelete(req *types.PathIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
