// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin_resource

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

type ResourceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceUpdateLogic {
	return &ResourceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceUpdateLogic) ResourceUpdate(req *types.ResourceUpdateReq) error {
	resource, err := l.svcCtx.ResourceModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("find resource by id: %d failed, err: %v", req.Id, err)
		return &exception.ResourceNotFound
	}
	copier.Copy(resource, req)
	resource.UpdatedAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.ResourceModel.Update(l.ctx, resource); err != nil {
		logx.Errorf("update resource by id: %d failed, err: %v", req.Id, err)
		return &exception.ResourceUpdateFailed
	}

	return nil
}
