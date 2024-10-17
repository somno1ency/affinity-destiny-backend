// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package category

import (
	"context"
	"time"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/category"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryCreateLogic {
	return &CategoryCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryCreateLogic) CategoryCreate(req *types.CategoryCreateReq) error {
	category := &category.Category{}
	copier.Copy(category, req)
	// TODO: set ownerId from token, assume ownerId = 1
	var ownerId int64 = 1
	category.OwnerId = ownerId
	category.CreatedAt = util.ConvertTime(time.Now())
	if _, err := l.svcCtx.CategoryModel.Insert(l.ctx, category); err != nil {
		logx.Errorf("insert category failed, err: %v", err)
		return &exception.CategoryCreateFailed
	}

	return nil
}
