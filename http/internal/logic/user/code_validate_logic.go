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
	"ad.com/pkg/repo/user"
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
	var currentUser *user.User
	currentUser, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil {
		// register this user when not exist
		currentUser = &user.User{
			CustomId:  util.GenCode(10, false),
			Mobile:    req.Mobile,
			Nickname:  req.Mobile,
			CreatedAt: util.ConvertTime(time.Now()),
		}
		if _, err := l.svcCtx.UserModel.Insert(l.ctx, currentUser); err != nil {
			logx.Errorf("insert user failed, err: %v", err)
			return nil, &exception.UserRegisterFailed
		}
		currentUser, _ = l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	}

	resp = &types.UserResp{}
	copier.Copy(resp, currentUser)
	resp.CreatedAt = currentUser.CreatedAt.Time.Format(shared.TimeFormatTemplate)
	resp.UpdatedAt = currentUser.UpdatedAt.Time.Format(shared.TimeFormatTemplate)
	resp.LastLoginAt = currentUser.LastLoginAt.Time.Format(shared.TimeFormatTemplate)

	currentUser.LastLoginAt = util.ConvertTime(time.Now())
	if err := l.svcCtx.UserModel.Update(l.ctx, currentUser); err != nil {
		logx.Errorf("update user login time failed, err: %v", err)
		return nil, &exception.UserLoginTimeUpdateFailed
	}

	return resp, nil
}
