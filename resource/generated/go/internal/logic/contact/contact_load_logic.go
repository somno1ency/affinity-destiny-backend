package contact

import (
	"context"

	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContactLoadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContactLoadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContactLoadLogic {
	return &ContactLoadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContactLoadLogic) ContactLoad(req *types.PathIdReq) (resp []types.UserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
