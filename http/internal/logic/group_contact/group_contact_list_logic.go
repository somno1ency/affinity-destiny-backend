// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/group_contact"
	"ad.com/pkg/shared"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactListLogic {
	return &GroupContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactListLogic) GroupContactList(req *types.GroupIdReq) (resp []types.GroupContactResp, err error) {
	groupContacts, err := l.svcCtx.GroupContactModel.Find(l.ctx, req.GroupId)
	if err != nil {
		logx.Errorf("find group contact list failed, err: %v", err)
		return nil, &exception.GroupContactListQueryFailed
	}
	groupContactsResp := util.TransformList(groupContacts, toResp)
	for i, groupContactResp := range groupContactsResp {
		var categoryResp = &types.CategoryResp{}
		if category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, groupContactResp.CategoryId); err == nil {
			copier.Copy(categoryResp, category)
		}
		var userResp = &types.UserResp{}
		if user, err := l.svcCtx.UserModel.FindOne(l.ctx, groupContactResp.UserId); err == nil {
			copier.Copy(userResp, user)
		}
		var groupResp = &types.GroupResp{}
		if group, err := l.svcCtx.UserModel.FindOne(l.ctx, groupContactResp.GroupId); err == nil {
			copier.Copy(groupResp, group)
		}
		groupContactsResp[i].Category = *categoryResp
		groupContactsResp[i].User = *userResp
		groupContactsResp[i].Group = *groupResp
	}

	return groupContactsResp, nil
}

func toResp(data *group_contact.GroupContact) types.GroupContactResp {
	return types.GroupContactResp{
		Id:             data.Id,
		GroupId:        data.GroupId.Int64,
		UserId:         data.UserId.Int64,
		CategoryId:     data.CategoryId,
		UserNickname:   data.UserNickname,
		Remark:         data.Remark,
		Background:     data.Background,
		IsDisturb:      data.IsDisturb,
		IsTop:          data.IsTop,
		IsShowNickname: data.IsShowNickname,
		ApprovalStatus: data.ApprovalStatus,
		ApprovalAt:     data.ApprovalAt.Time.Format(shared.TimeFormatTemplate),
		CreatedAt:      data.CreatedAt.Time.Format(shared.TimeFormatTemplate),
		UpdatedAt:      data.UpdatedAt.Time.Format(shared.TimeFormatTemplate),
	}
}
