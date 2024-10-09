package user_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserContactListLogic {
	return &UserContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserContactListLogic) UserContactList() (resp []types.UserContactResp, err error) {
	// todo: add your logic here and delete this line

	return
}
