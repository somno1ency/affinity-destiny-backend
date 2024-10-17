// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package category

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryDeleteLogic {
	return &CategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryDeleteLogic) CategoryDelete(req *types.PathIdReq) error {
	if _, err := l.svcCtx.CategoryModel.FindOne(l.ctx, req.Id); err != nil {
		logx.Errorf("find category by id: %d failed, err: %v", req.Id, err)
		return &exception.CategoryNotFound
	}
	if err := l.svcCtx.CategoryModel.Delete(l.ctx, req.Id); err != nil {
		logx.Errorf("delete category by id: %d failed, err: %v", req.Id, err)
		return &exception.CategoryDeleteFailed
	}

	return nil
}
