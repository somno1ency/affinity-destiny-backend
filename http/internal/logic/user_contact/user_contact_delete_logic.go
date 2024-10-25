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
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	// TODO: assume delete user contact for user 1
	var ownerId int64 = 1
	if ownerId == req.Id {
		return &exception.UserContactDeleteFailed
	}
	toTarget, err := l.svcCtx.UserContactModel.FindDstContact(l.ctx, ownerId, req.Id)
	if err != nil {
		logx.Errorf("find self contact failed, err: %v", err)
		return &exception.UserContactTargetNotFriend
	}
	toSelf, err := l.svcCtx.UserContactModel.FindDstContact(l.ctx, req.Id, ownerId)
	if err != nil {
		logx.Errorf("find target contact failed, err: %v", err)
		return &exception.UserContactNotTargetFriend
	}

	toTarget.ReApply = 1
	toTarget.UpdatedAt = util.ConvertTime(time.Now())
	toSelf.ReApply = 1
	toSelf.UpdatedAt = util.ConvertTime(time.Now())
	l.svcCtx.Conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.UserContactModel.Update(ctx, toTarget); err != nil {
			return &exception.UserContactUpdateFailed
		}
		if err := l.svcCtx.UserContactModel.Update(ctx, toSelf); err != nil {
			return &exception.UserContactUpdateFailed
		}

		return nil
	})

	return nil
}
