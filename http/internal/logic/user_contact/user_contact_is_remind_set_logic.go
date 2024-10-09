package user_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactIsRemindSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserContactIsRemindSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserContactIsRemindSetLogic {
	return &UserContactIsRemindSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserContactIsRemindSetLogic) UserContactIsRemindSet(req *types.UserContactIsSet) error {
	// todo: add your logic here and delete this line

	return nil
}