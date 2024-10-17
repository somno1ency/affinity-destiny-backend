// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/user_contact"
	"ad.com/pkg/shared"
	shared_types "ad.com/pkg/shared/types"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserContactListLogic {
	return &UserContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserContactListLogic) UserContactList(req *types.QueryListReq) (resp *types.UserContactPagedResp, err error) {
	modelReq := &shared_types.QueryListReq{}
	copier.Copy(modelReq, req)
	// TODO: here need transfer ownerId from token to backend, assume ownerId = 1
	var ownerId int64 = 1
	userContacts, err := l.svcCtx.UserContactModel.Find(l.ctx, ownerId, modelReq)
	if err != nil {
		logx.Errorf("find user contact list failed, err: %v", err)
		return nil, &exception.UserContactListQueryFailed
	}
	count, err := l.svcCtx.GroupModel.Count(l.ctx, ownerId)
	if err != nil {
		logx.Errorf("count user contact list failed, err: %v", err)
		return nil, &exception.GroupListQueryFailed
	}

	userContactsResp := util.TransformList(userContacts, toResp)
	for i, userContactResp := range userContactsResp {
		var categoryResp = &types.CategoryResp{}
		if category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, userContactResp.CategoryId); err == nil {
			copier.Copy(categoryResp, category)
		}
		var userResp = &types.UserResp{}
		if user, err := l.svcCtx.UserModel.FindOne(l.ctx, userContactResp.DstId); err == nil {
			copier.Copy(userResp, user)
		}
		userContactsResp[i].Category = *categoryResp
		userContactsResp[i].User = *userResp
	}

	resp = &types.UserContactPagedResp{}
	resp.Data = userContactsResp
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.TotalCount = count
	resp.TotalPage = (count / req.PageSize) + 1

	return resp, nil
}

func toResp(data *user_contact.UserContact) types.UserContactResp {

	return types.UserContactResp{
		Id:         data.Id,
		OwnerId:    data.OwnerId,
		DstId:      data.DstId,
		CategoryId: data.CategoryId,
		Background: data.Background,
		IsDisturb:  data.IsDisturb,
		IsTop:      data.IsTop,
		IsRemind:   data.IsRemind,
		CreatedAt:  data.CreatedAt.Time.Format(shared.TimeFormatTemplate),
		UpdatedAt:  data.UpdatedAt.Time.Format(shared.TimeFormatTemplate),
	}
}
