package user_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactCategorySetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserContactCategorySetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserContactCategorySetLogic {
	return &UserContactCategorySetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserContactCategorySetLogic) UserContactCategorySet(req *types.UserContactCategorySetReq) error {
	// todo: add your logic here and delete this line

	return nil
}
