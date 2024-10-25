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
	"ad.com/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactNicknameSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactNicknameSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactNicknameSetLogic {
	return &GroupContactNicknameSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactNicknameSetLogic) GroupContactNicknameSet(req *types.GroupContactStringSetReq) error {
	// TODO: assume operate for user 1
	var ownerId int64 = 1
	groupContact, err := l.svcCtx.GroupContactModel.FindByUserId(l.ctx, req.Id, ownerId)
	if err != nil {
		logx.Errorf("find group contact failed, err: %v", err)
		return &exception.GroupContactNotFound
	}
	groupContact.UserNickname = req.Value
	groupContact.UpdatedAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.GroupContactModel.Update(l.ctx, groupContact); err != nil {
		logx.Errorf("update user contact failed, err: %v", err)
		return &exception.GroupContactUpdateFailed
	}

	return nil
}
