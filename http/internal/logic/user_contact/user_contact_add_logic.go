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
	"ad.com/pkg/repo/user_contact"
	"ad.com/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("find user failed, err: %v", err)
		return &exception.UserNotFound
	}
	// TODO: assume add user contact for user 1
	var ownerId int64 = 1
	if ownerId == user.Id {
		return &exception.UserContactAddSelfFailed
	}
	if _, err := l.svcCtx.UserContactModel.FindDstContact(l.ctx, ownerId, req.Id); err == nil {
		return &exception.UserContactAlreadyExist
	}

	selfContact := &user_contact.UserContact{}
	selfContact.OwnerId = ownerId
	selfContact.DstId = user.Id
	selfContact.IsInitiator = 1
	selfContact.CreatedAt = util.ConvertTime(time.Now())

	targetContact := &user_contact.UserContact{}
	targetContact.OwnerId = user.Id
	targetContact.DstId = ownerId
	targetContact.IsInitiator = 0
	targetContact.CreatedAt = util.ConvertTime(time.Now())
	l.svcCtx.Conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if _, err := l.svcCtx.UserContactModel.Insert(ctx, selfContact); err != nil {
			logx.Errorf("insert self contact failed, err: %v", err)
			return &exception.UserContactCreateFailed
		}
		if _, err := l.svcCtx.UserContactModel.Insert(ctx, targetContact); err != nil {
			logx.Errorf("insert target contact failed, err: %v", err)
			return &exception.UserContactCreateFailed
		}

		return nil
	})

	return nil
}
