// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"context"
	"fmt"
	"time"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/shared"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CodeValidateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCodeValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CodeValidateLogic {
	return &CodeValidateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CodeValidateLogic) CodeValidate(req *types.CodeValidateReq) (resp *types.UserResp, err error) {
	code, err := l.svcCtx.RdsClient.Get(fmt.Sprintf(shared.CodeSendToPhone, req.Mobile))
	if err != nil || code != req.Code {
		logx.Errorf("validate code err, current code: %s, target code: %s", req.Code, code)
		return nil, &exception.CodeValidateFailed
	}
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil {
		logx.Errorf("find user by mobile failed, err: %v", err)
		return nil, &exception.UserNotFound
	}

	resp = &types.UserResp{}
	copier.Copy(resp, user)
	resp.CreatedAt = user.CreatedAt.Time.Format(shared.TimeFormatTemplate)
	resp.UpdatedAt = user.UpdatedAt.Time.Format(shared.TimeFormatTemplate)
	resp.LastLoginAt = user.LastLoginAt.Time.Format(shared.TimeFormatTemplate)

	user.LastLoginAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
		logx.Errorf("update user login time failed, err: %v", err)
		return nil, &exception.UserLoginTimeUpdateFailed
	}

	return resp, nil
}
