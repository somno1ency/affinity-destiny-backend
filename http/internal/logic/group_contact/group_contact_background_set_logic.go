package group_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactBackgroundSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactBackgroundSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactBackgroundSetLogic {
	return &GroupContactBackgroundSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactBackgroundSetLogic) GroupContactBackgroundSet(req *types.GroupContactStringSetReq) error {
	// todo: add your logic here and delete this line

	return nil
}
