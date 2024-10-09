package user_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserContactDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserContactDeleteLogic {
	return &UserContactDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserContactDeleteLogic) UserContactDelete(req *types.PathIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}