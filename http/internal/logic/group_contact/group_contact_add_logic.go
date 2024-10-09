package group_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactAddLogic {
	return &GroupContactAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactAddLogic) GroupContactAdd(req *types.GroupContactAddReq) error {
	// todo: add your logic here and delete this line

	return nil
}
