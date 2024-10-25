// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group_contact

import (
	"context"
	"time"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/group_contact"
	"ad.com/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactJoinLogic {
	return &GroupContactJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactJoinLogic) GroupContactJoin(req *types.GroupIdReq) error {
	// TODO: assume operate for user 1
	var ownerId int64 = 1
	if _, err := l.svcCtx.GroupContactModel.FindByUserId(l.ctx, req.GroupId, ownerId); err == nil {
		logx.Errorf("group contact already exists, err: %v", err)
		return &exception.GroupContactAlreadyExists
	}
	groupContact := &group_contact.GroupContact{}
	groupContact.GroupId = util.ConvertInt64(req.GroupId)
	groupContact.UserId = util.ConvertInt64(ownerId)
	groupContact.CreatedAt = util.ConvertTime(time.Now())
	if _, err := l.svcCtx.GroupContactModel.Insert(l.ctx, groupContact); err != nil {
		logx.Errorf("join group contact failed, err: %v", err)
		return &exception.GroupContactJoinFailed
	}

	return nil
}
