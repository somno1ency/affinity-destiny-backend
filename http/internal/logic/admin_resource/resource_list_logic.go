// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin_resource

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/resource"
	"ad.com/pkg/shared"
	shared_types "ad.com/pkg/shared/types"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceListLogic {
	return &ResourceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceListLogic) ResourceList(req *types.ResourceListReq) (resp *types.ResourcePagedResp, err error) {
	modelReq := &shared_types.QueryListReq{}
	copier.Copy(modelReq, req)
	resources, err := l.svcCtx.ResourceModel.Find(l.ctx, req.Type, modelReq)
	if err != nil {
		logx.Errorf("find resource list failed, err: %v", err)
		return nil, &exception.ResourceListQueryFailed
	}
	count, err := l.svcCtx.ResourceModel.Count(l.ctx, req.Type)
	if err != nil {
		logx.Errorf("count resource list failed, err: %v", err)
		return nil, &exception.ResourceListQueryFailed
	}

	resp = &types.ResourcePagedResp{}
	resp.Data = util.TransformList(resources, toResp)
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.TotalCount = count
	resp.TotalPage = (count / req.PageSize) + 1

	return resp, nil
}

func toResp(data *resource.Resource) types.ResourceResp {
	return types.ResourceResp{
		Id:        data.Id,
		Src:       data.Src,
		Type:      data.Type,
		NameZh:    data.NameZh,
		NameEn:    data.NameEn,
		CreatedAt: data.CreatedAt.Time.Format(shared.TimeFormatTemplate),
		UpdatedAt: data.UpdatedAt.Time.Format(shared.TimeFormatTemplate),
	}
}
