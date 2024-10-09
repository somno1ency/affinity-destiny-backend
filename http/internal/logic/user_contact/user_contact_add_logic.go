package user_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserContactAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserContactAddLogic {
	return &UserContactAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserContactAddLogic) UserContactAdd(req *types.PathIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
