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

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateReq) error {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("find user by id: %d failed, err: %v", req.Id, err)
		return &exception.UserNotFound
	}
	user.CustomId = req.CustomId
	user.Nickname = req.Nickname
	user.Avatar = req.Avatar
	user.Sex = req.Sex
	user.Memo = req.Memo
	user.UpdatedAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
		logx.Errorf("update user by id: %d failed, err: %v", req.Id, err)
		return &exception.UserUpdateFailed
	}

	return nil
}
