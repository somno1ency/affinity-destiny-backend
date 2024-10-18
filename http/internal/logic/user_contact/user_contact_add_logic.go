// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserContactAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserContactAddLogic {
	return &UserContactAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserContactAddLogic) UserContactAdd(req *types.PathIdReq) error {
	// TODO: assume add user contact for user 1
	// var ownerId int64 = 1
	// if _, err := l.svcCtx.UserContactModel.FindDstContact(l.ctx, ownerId, req.Id); err == nil {
	// 	return &exception.UserContactAlreadyExists
	// }
	// group.OwnerId = ownerId
	// group.CustomId = util.GenCode(10, false)
	// group.CreatedAt = util.ConvertTime(time.Now())
	// if _, err := l.svcCtx.GroupModel.Insert(l.ctx, group); err != nil {
	// 	logx.Errorf("insert group failed, err: %v", err)
	// 	return &exception.GroupCreateFailed
	// }

	return nil
}
