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
	"github.com/zeromicro/x/errors"
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
	if toTarget, err := l.svcCtx.UserContactModel.FindDstContact(l.ctx, ownerId, req.Id); err == nil {
		switch toTarget.ApprovalStatus {
		case 0:
			// wait for process
			if toTarget.IsInitiator == 1 {
				// invite by self, target don't process
				return &exception.UserContactSelfInvite
			} else {
				// invite by target, self don't process
				return &exception.UserContactTargetInvite
			}
		case 10:
			// already process
			// no matter who invite, they are already friends
			return &exception.UserContactAlreadyExist
		case 20:
			// denied
			// no matter who denied, they can re-invite, so make ReApply = 1 of the two records
			toSelf, err := l.svcCtx.UserContactModel.FindDstContact(l.ctx, req.Id, ownerId)
			if err != nil {
				logx.Errorf("find toSelf contact failed, err: %v", err)
				return errors.New(exception.UnknownError.Code, err.Error())
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
		}
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
