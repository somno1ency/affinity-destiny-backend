package group

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDeleteLogic {
	return &GroupDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupDeleteLogic) GroupDelete(req *types.PathIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
