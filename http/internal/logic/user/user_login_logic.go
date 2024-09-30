package user

import (
	"context"
	"fmt"
	"time"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, util.ConvertString(req.Mobile))
	if err != nil {
		return nil, errors.New(exception.UserNotFound.Code, exception.UserNotFound.Msg)
	}
	if !util.ValidatePassword(req.Password, user.Salt.String, user.Password.String) {
		return nil, errors.New(exception.UserNotMatch.Code, exception.UserNotMatch.Msg)
	}

	resp = &types.UserResp{}
	str := util.Md5Encode((fmt.Sprintf("%d", time.Now().Unix())))
	user.Token = util.ConvertString(str)
	user.LastLoginAt = util.ConvertTime(time.Now())
	_ = l.svcCtx.UserModel.Update(l.ctx, user)
	_ = copier.Copy(resp, user)

	return resp, nil
}
