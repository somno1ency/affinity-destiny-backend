// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group

import (
	"context"
	"time"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/group"
	"ad.com/pkg/repo/group_contact"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GroupCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCreateLogic) GroupCreate(req *types.GroupCreateReq) error {
	group := &group.Group{}
	copier.Copy(group, req)
	// TODO: set ownerId from token, assume ownerId = 1
	var ownerId int64 = 1
	group.OwnerId = ownerId

	var customId string
	for {
		customId = util.GenCode(10, false)
		if _, err := l.svcCtx.GroupModel.FindOneByCustomId(l.ctx, customId); err != nil {
			break
		}
		logx.Infof("duplicate customId: %s, try again", customId)
	}
	group.CustomId = customId
	group.CreatedAt = util.ConvertTime(time.Now())
	l.svcCtx.Conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		insertResult, err := l.svcCtx.GroupModel.Insert(l.ctx, group)
		if err != nil {
			logx.Errorf("insert group failed, err: %v", err)
			return &exception.GroupCreateFailed
		}
		insertId, err := insertResult.LastInsertId()
		if err != nil {
			logx.Errorf("get insertId failed, err: %v", err)
			return &exception.GroupCreateFailed
		}

		// create group contact for owner
		groupContact := &group_contact.GroupContact{}
		groupContact.GroupId = util.ConvertInt64(insertId)
		groupContact.UserId = util.ConvertInt64(ownerId)
		groupContact.CreatedAt = util.ConvertTime(time.Now())
		if _, err := l.svcCtx.GroupContactModel.Insert(ctx, groupContact); err != nil {
			logx.Errorf("insert group contact for ownerId: %d, groupId: %d, failed, err: %v", ownerId, insertId, err)
			return &exception.UserContactCreateFailed
		}

		return nil
	})

	return nil
}
