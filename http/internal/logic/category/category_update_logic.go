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
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryUpdateLogic {
	return &CategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryUpdateLogic) CategoryUpdate(req *types.CategoryUpdateReq) error {
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("find category by id: %d failed, err: %v", req.Id, err)
		return &exception.CategoryNotFound
	}
	copier.Copy(category, req)
	category.UpdatedAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.CategoryModel.Update(l.ctx, category); err != nil {
		logx.Errorf("update category by id: %d failed, err: %v", req.Id, err)
		return &exception.CategoryUpdateFailed
	}

	return nil
}
