// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"context"
	"time"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/shared"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil {
		logx.Errorf("find user by mobile failed, err: %v", err)
		return nil, &exception.UserNotFound
	}
	if !util.ValidatePassword(req.Password, user.Salt, user.Password) {
		logx.Errorf("validate password failed for user: %s", user.Mobile)
		return nil, &exception.UserPasswordValidateFailed
	}

	resp = &types.UserResp{}
	copier.Copy(resp, user)
	resp.CreatedAt = user.CreatedAt.Time.Format(shared.TimeFormatTemplate)
	resp.UpdatedAt = user.UpdatedAt.Time.Format(shared.TimeFormatTemplate)
	resp.LastLoginAt = user.LastLoginAt.Time.Format(shared.TimeFormatTemplate)

	user.LastLoginAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
		logx.Errorf("update user login time failed, err: %v", err)
		return nil, &exception.UserLoginTimeUpdateFailed
	}

	return resp, nil
}
