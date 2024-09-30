package user

import (
	"context"
	"fmt"
	"math/rand"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/user"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/x/errors"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserResp, err error) {
	_, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, util.ConvertString(req.Mobile))
	if err == nil {
		return nil, errors.New(exception.UserExists.Code, exception.UserExists.Msg)
	}
	if err != sqlx.ErrNotFound {
		return nil, errors.New(exception.UnknownError.Code, err.Error())
	}

	user := &user.User{}
	_ = copier.Copy(user, req)
	// 处理前端未携带的其他字段
	user.Salt = util.ConvertString(fmt.Sprintf("%06d", rand.Int31n(10000)))
	user.Password = util.ConvertString(util.MakePassword(req.Password, user.Salt.String))
	user.Token = util.ConvertString(fmt.Sprintf("%08d", rand.Int31()))
	// 时间会自动处理
	// user.CreatedAt = util.ConvertTime(time.Now())
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, errors.New(exception.UserRegisterFailed.Code, exception.UserRegisterFailed.Msg)
	}

	resp = &types.UserResp{}
	_ = copier.Copy(resp, user)

	return resp, nil
}
