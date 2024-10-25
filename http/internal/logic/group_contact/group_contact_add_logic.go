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

type GroupContactAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactAddLogic {
	return &GroupContactAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactAddLogic) GroupContactAdd(req *types.GroupContactAddReq) error {
	if len(req.UserIds) == 0 {
		return &exception.GroupContactUserIdsEmpty
	}
	if _, err := l.svcCtx.GroupModel.FindOne(l.ctx, req.GroupId); err != nil {
		logx.Errorf("find group failed, err: %v", err)
		return &exception.GroupNotFound
	}
	users, err := l.svcCtx.UserModel.FindByUserIds(l.ctx, req.UserIds)
	if err != nil {
		logx.Errorf("find users failed, err: %v", err)
		return &exception.UserNotFound
	}
	notExistIds := make([]int64, 0)
	userIds := make([]int64, 0)
	for _, user := range users {
		userIds = append(userIds, user.Id)
	}
	for _, userId := range req.UserIds {
		if !util.Contains(userIds, userId) {
			notExistIds = append(notExistIds, userId)
		}
	}
	if len(notExistIds) > 0 {
		logx.Infof("groupId: %d, request userIds: %v, not exist userIds: %v", req.GroupId, req.UserIds, notExistIds)
	}
	if len(notExistIds) == len(req.UserIds) {
		logx.Infof("groupId: %d, request userIds: %v, all user are not exist...", req.GroupId, req.UserIds)
		return &exception.GroupContactAllUserNotExist
	}
	var groupContacts []*group_contact.GroupContact
	for _, user := range users {
		if _, err := l.svcCtx.GroupContactModel.FindByUserId(l.ctx, req.GroupId, user.Id); err == nil {
			continue
		}
		groupContact := &group_contact.GroupContact{}
		groupContact.GroupId = util.ConvertInt64(req.GroupId)
		groupContact.UserId = util.ConvertInt64(user.Id)
		groupContact.CreatedAt = util.ConvertTime(time.Now())
		groupContacts = append(groupContacts, groupContact)
	}
	if err = l.svcCtx.GroupContactModel.BatchInsert(l.ctx, groupContacts); err != nil {
		logx.Errorf("batch insert group contact failed, err: %v", err)
		return &exception.GroupContactsCreateFailed
	}

	return nil
}
