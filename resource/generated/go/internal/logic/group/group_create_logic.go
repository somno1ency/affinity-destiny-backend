package group

import (
	"context"

	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCreateLogic) GroupCreate(req *types.GroupCreateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
