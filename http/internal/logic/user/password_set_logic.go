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
	"ad.com/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type PasswordSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPasswordSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordSetLogic {
	return &PasswordSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PasswordSetLogic) PasswordSet(req *types.UserPasswordReq) error {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("find user by id: %d failed, err: %v", req.Id, err)
		return &exception.UserNotFound
	}
	canUpdatePassword := false
	if user.Password == "" {
		canUpdatePassword = true
	} else {
		if !util.ValidatePassword(req.OldPassword, user.Salt, user.Password) {
			logx.Errorf("validate password failed for user: %s", user.Mobile)
			return &exception.UserPasswordValidateFailed
		}
		canUpdatePassword = true
	}
	if canUpdatePassword {
		user.Salt = util.GenCode(8, true)
		user.Password = util.MakePassword(req.NewPassword, user.Salt)
		user.UpdatedAt = util.ConvertTime(time.Now())
		if err := l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
			logx.Errorf("update user password by id: %d failed, err: %v", req.Id, err)
			return &exception.UserPasswordUpdateFailed
		}
	}

	return nil
}
