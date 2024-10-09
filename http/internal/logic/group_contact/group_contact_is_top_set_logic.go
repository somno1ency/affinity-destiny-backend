package group_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactIsTopSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactIsTopSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactIsTopSetLogic {
	return &GroupContactIsTopSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactIsTopSetLogic) GroupContactIsTopSet(req *types.GroupContactIsSet) error {
	// todo: add your logic here and delete this line

	return nil
}
