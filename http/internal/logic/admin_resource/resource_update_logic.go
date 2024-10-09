package admin_resource

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceUpdateLogic {
	return &ResourceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceUpdateLogic) ResourceUpdate(req *types.ResourceUpdateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
