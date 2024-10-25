// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user_contact

import (
	"context"
	"time"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactIsTopSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserContactIsTopSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserContactIsTopSetLogic {
	return &UserContactIsTopSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserContactIsTopSetLogic) UserContactIsTopSet(req *types.UserContactIsSet) error {
	// TODO: assume operate for user 1
	var ownerId int64 = 1
	userContact, err := l.svcCtx.UserContactModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("find user contact failed, err: %v", err)
		return &exception.UserContactTargetNotFriend
	}
	if ownerId != userContact.OwnerId {
		return &exception.NoPermissionError
	}
	var isTop int64 = 0
	if req.Value {
		isTop = 1
	}
	userContact.IsTop = isTop
	userContact.UpdatedAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.UserContactModel.Update(l.ctx, userContact); err != nil {
		logx.Errorf("update user contact failed, err: %v", err)
		return &exception.UserContactUpdateFailed
	}

	return nil
}
