// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupDisbandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupDisbandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDisbandLogic {
	return &GroupDisbandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupDisbandLogic) GroupDisband(req *types.PathIdReq) error {
	if _, err := l.svcCtx.GroupModel.FindOne(l.ctx, req.Id); err != nil {
		logx.Errorf("find group by id: %d failed, err: %v", req.Id, err)
		return &exception.GroupNotFound
	}
	if err := l.svcCtx.GroupModel.Delete(l.ctx, req.Id); err != nil {
		logx.Errorf("disband group by id: %d failed, err: %v", req.Id, err)
		return &exception.GroupDisbandFailed
	}

	return nil
}
