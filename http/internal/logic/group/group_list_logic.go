// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/group"
	"ad.com/pkg/shared"
	shared_types "ad.com/pkg/shared/types"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList(req *types.QueryListReq) (resp *types.GroupPagedResp, err error) {
	modelReq := &shared_types.QueryListReq{}
	copier.Copy(modelReq, req)
	// TODO: here need transfer ownerId from token to backend, assume ownerId = 1
	var ownerId int64 = 1
	groups, err := l.svcCtx.GroupModel.Find(l.ctx, ownerId, modelReq)
	if err != nil {
		logx.Errorf("find group list failed, err: %v", err)
		return nil, &exception.GroupListQueryFailed
	}
	count, err := l.svcCtx.GroupModel.Count(l.ctx, ownerId)
	if err != nil {
		logx.Errorf("count group list failed, err: %v", err)
		return nil, &exception.GroupListQueryFailed
	}

	resp = &types.GroupPagedResp{}
	resp.Data = util.TransformList(groups, toResp)
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.TotalCount = count
	resp.TotalPage = (count / req.PageSize) + 1

	return resp, nil
}

func toResp(data *group.Group) types.GroupResp {
	return types.GroupResp{
		Id:        data.Id,
		CustomId:  data.CustomId,
		Name:      data.Name,
		OwnerId:   data.OwnerId,
		Avatar:    data.Avatar,
		Memo:      data.Memo,
		CreatedAt: data.CreatedAt.Time.Format(shared.TimeFormatTemplate),
		UpdatedAt: data.UpdatedAt.Time.Format(shared.TimeFormatTemplate),
	}
}
