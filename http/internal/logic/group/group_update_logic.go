// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group

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

type GroupUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUpdateLogic {
	return &GroupUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUpdateLogic) GroupUpdate(req *types.GroupUpdateReq) error {
	group, err := l.svcCtx.GroupModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("find group by id: %d failed, err: %v", req.Id, err)
		return &exception.GroupNotFound
	}
	copier.Copy(group, req)
	group.UpdatedAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.GroupModel.Update(l.ctx, group); err != nil {
		logx.Errorf("update group by id: %d failed, err: %v", req.Id, err)
		return &exception.GroupUpdateFailed
	}

	return nil
}
