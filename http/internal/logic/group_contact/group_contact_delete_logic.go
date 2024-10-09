package group_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactDeleteLogic {
	return &GroupContactDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactDeleteLogic) GroupContactDelete(req *types.PathIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
