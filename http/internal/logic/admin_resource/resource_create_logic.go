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
	"ad.com/pkg/repo/resource"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ResourceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceCreateLogic {
	return &ResourceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceCreateLogic) ResourceCreate(req *types.ResourceCreateReq) error {
	resource := &resource.Resource{}
	copier.Copy(resource, req)
	resource.CreatedAt = util.ConvertTime(time.Now())
	if _, err := l.svcCtx.ResourceModel.Insert(l.ctx, resource); err != nil {
		logx.Errorf("insert resource failed, err: %v", err)
		return &exception.ResourceCreateFailed
	}

	return nil
}
