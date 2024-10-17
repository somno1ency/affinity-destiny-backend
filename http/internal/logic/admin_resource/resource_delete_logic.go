// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin_resource

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceDeleteLogic {
	return &ResourceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceDeleteLogic) ResourceDelete(req *types.PathIdReq) error {
	if _, err := l.svcCtx.ResourceModel.FindOne(l.ctx, req.Id); err != nil {
		logx.Errorf("find resource by id: %d failed, err: %v", req.Id, err)
		return &exception.ResourceNotFound
	}
	if err := l.svcCtx.ResourceModel.Delete(l.ctx, req.Id); err != nil {
		logx.Errorf("delete resource by id: %d failed, err: %v", req.Id, err)
		return &exception.ResourceDeleteFailed
	}

	return nil
}
