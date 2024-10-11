package user

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/shared"

	"github.com/jinzhu/copier"
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
	resp = &types.UserResp{}
	user, err := l.svcCtx.UserModel.FindOneByCustomId(l.ctx, req.Id)
	if err != nil {
		user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Id)
		if err != nil {
			logx.Errorf("find user by custom id: %s failed, err: %v", req.Id, err)
			return nil, &exception.UserNotFound
		}
	}
	copier.Copy(resp, user)
	resp.CreatedAt = user.CreatedAt.Time.Format(shared.TimeFormatTemplate)
	resp.UpdatedAt = user.UpdatedAt.Time.Format(shared.TimeFormatTemplate)
	resp.LastLoginAt = user.LastLoginAt.Time.Format(shared.TimeFormatTemplate)

	return resp, nil
}
