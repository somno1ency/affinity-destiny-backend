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

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupTransferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupTransferLogic {
	return &GroupTransferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupTransferLogic) GroupTransfer(req *types.GroupTransferReq) error {
	group, err := l.svcCtx.GroupModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("find group by id: %d failed, err: %v", req.Id, err)
		return &exception.GroupNotFound
	}
	if group.OwnerId == req.OwnerId {
		logx.Errorf("can't transfer group to yourself, group id: %d", req.Id)
		return &exception.GroupTransferToSelfError
	}
	group.OwnerId = req.OwnerId
	group.UpdatedAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.GroupModel.Update(l.ctx, group); err != nil {
		logx.Errorf("transfer group by id: %d failed, err: %v", req.Id, err)
		return &exception.GroupTransferFailed
	}

	return nil
}
