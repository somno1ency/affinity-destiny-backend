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
	"ad.com/pkg/repo/group"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCreateLogic) GroupCreate(req *types.GroupCreateReq) error {
	group := &group.Group{}
	copier.Copy(group, req)
	// TODO: set ownerId from token, assume ownerId = 1
	var ownerId int64 = 1
	group.OwnerId = ownerId

	var customId string
	for {
		customId = util.GenCode(10, false)
		if _, err := l.svcCtx.GroupModel.FindOneByCustomId(l.ctx, customId); err != nil {
			break
		}
		logx.Infof("duplicate customId: %s, try again", customId)
	}
	group.CustomId = customId
	group.CreatedAt = util.ConvertTime(time.Now())
	if _, err := l.svcCtx.GroupModel.Insert(l.ctx, group); err != nil {
		logx.Errorf("insert group failed, err: %v", err)
		return &exception.GroupCreateFailed
	}

	return nil
}
