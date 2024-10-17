// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package category

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/category"
	"ad.com/pkg/shared"
	shared_types "ad.com/pkg/shared/types"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList(req *types.QueryListReq) (resp *types.CategoryPagedResp, err error) {
	modelReq := &shared_types.QueryListReq{}
	copier.Copy(modelReq, req)
	// TODO: here need transfer ownerId from token to backend, assume ownerId = 1
	var ownerId int64 = 1
	categories, err := l.svcCtx.CategoryModel.Find(l.ctx, ownerId, modelReq)
	if err != nil {
		logx.Errorf("find category list failed, err: %v", err)
		return nil, &exception.CategoryListQueryFailed
	}
	count, err := l.svcCtx.CategoryModel.Count(l.ctx, ownerId)
	if err != nil {
		logx.Errorf("count category list failed, err: %v", err)
		return nil, &exception.CategoryListQueryFailed
	}

	resp = &types.CategoryPagedResp{}
	resp.Data = util.TransformList(categories, toResp)
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.TotalCount = count
	resp.TotalPage = (count / req.PageSize) + 1

	return resp, nil
}

func toResp(data *category.Category) types.CategoryResp {
	return types.CategoryResp{
		Id:        data.Id,
		OwnerId:   data.OwnerId,
		NameZh:    data.NameZh,
		NameEn:    data.NameEn,
		Star:      data.Star,
		CreatedAt: data.CreatedAt.Time.Format(shared.TimeFormatTemplate),
		UpdatedAt: data.UpdatedAt.Time.Format(shared.TimeFormatTemplate),
	}
}
