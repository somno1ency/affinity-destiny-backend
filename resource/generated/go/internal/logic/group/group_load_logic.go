package group

import (
	"context"

	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupLoadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupLoadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupLoadLogic {
	return &GroupLoadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupLoadLogic) GroupLoad(req *types.PathIdReq) (resp []types.GroupResp, err error) {
	// todo: add your logic here and delete this line

	return
}
