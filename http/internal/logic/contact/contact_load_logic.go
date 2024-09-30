package contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/jinzhu/copier"
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
	resp = make([]types.UserResp, 0)
	userIds, _ := l.svcCtx.UserContactModel.FindDstIdsByOwnerId(l.ctx, req.Id)
	if len(userIds) == 0 {
		return resp, nil
	}
	users, _ := l.svcCtx.UserModel.FindInIds(l.ctx, userIds)
	_ = copier.Copy(&resp, users)

	return resp, nil
}
