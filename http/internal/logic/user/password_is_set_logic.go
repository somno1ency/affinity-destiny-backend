package user

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"

	"github.com/zeromicro/go-zero/core/logx"
)

type PasswordIsSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPasswordIsSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordIsSetLogic {
	return &PasswordIsSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PasswordIsSetLogic) PasswordIsSet(req *types.PathIdReq) (resp *types.PasswordResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("find user by id: %d failed, err: %v", req.Id, err)
		return nil, &exception.UserNotFound
	}

	resp = &types.PasswordResp{}
	resp.IsSetPassword = user.Password != ""

	return resp, nil
}
