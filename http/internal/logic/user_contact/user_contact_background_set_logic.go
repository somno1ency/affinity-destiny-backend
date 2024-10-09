package user_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactBackgroundSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserContactBackgroundSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserContactBackgroundSetLogic {
	return &UserContactBackgroundSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserContactBackgroundSetLogic) UserContactBackgroundSet(req *types.UserContactBackgroundSetReq) error {
	// todo: add your logic here and delete this line

	return nil
}