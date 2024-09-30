package group

import (
	"context"

	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupJoinLogic {
	return &GroupJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupJoinLogic) GroupJoin(req *types.GroupJoinReq) error {
	// todo: add your logic here and delete this line

	return nil
}
