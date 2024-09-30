package user

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
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

func (l *UserFindLogic) UserFind(req *types.PathIdReq) (resp *types.UserResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errors.New(exception.UserNotFound.Code, exception.UserNotFound.Msg)
	}

	resp = &types.UserResp{}
	_ = copier.Copy(resp, user)

	return resp, nil
}
