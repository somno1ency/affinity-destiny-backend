// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactLeaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactLeaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactLeaveLogic {
	return &GroupContactLeaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactLeaveLogic) GroupContactLeave(req *types.GroupIdReq) error {
	// TODO: assume operate for user 1
	var ownerId int64 = 1
	groupContact, err := l.svcCtx.GroupContactModel.FindByUserId(l.ctx, req.GroupId, ownerId)
	if err != nil {
		logx.Errorf("find group contact failed, err: %v", err)
		return &exception.GroupContactNotFound
	}
	if err := l.svcCtx.GroupContactModel.Delete(l.ctx, groupContact.Id); err != nil {
		logx.Errorf("delete group contact failed, err: %v", err)
		return &exception.GroupContactDeleteFailed
	}

	return nil
}
