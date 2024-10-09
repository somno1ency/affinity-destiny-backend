package admin_resource

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceCreateLogic {
	return &ResourceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceCreateLogic) ResourceCreate(req *types.ResourceCreateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
