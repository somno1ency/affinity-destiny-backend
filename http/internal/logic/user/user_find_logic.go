package user

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFindLogic {
	return &UserFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFindLogic) UserFind(req *types.PathIdStrReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
