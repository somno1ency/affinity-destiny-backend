package group_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactJoinLogic {
	return &GroupContactJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactJoinLogic) GroupContactJoin(req *types.GroupIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}