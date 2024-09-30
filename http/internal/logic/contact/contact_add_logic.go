package contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	user_contact "ad.com/pkg/repo/user-contact"
	"ad.com/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/x/errors"
)

type ContactAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContactAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContactAddLogic {
	return &ContactAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContactAddLogic) ContactAdd(req *types.ContactAddReq) error {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, util.ConvertString(req.Mobile))
	if err != nil {
		if err == sqlx.ErrNotFound {
			return errors.New(exception.UserNotFound.Code, exception.UserNotFound.Msg)
		}

		return errors.New(exception.UserNotFound.Code, err.Error())
	}
	if user.Id == req.OwnerId {
		return errors.New(exception.UserAddSelfFailed.Code, exception.UserAddSelfFailed.Msg)
	}
	// 判断是否已经加了好友
	if isContact, _ := l.svcCtx.UserContactModel.IsContact(l.ctx, req.OwnerId, user.Id); isContact {
		return errors.New(exception.UserAlreadyContact.Code, exception.UserAlreadyContact.Msg)
	}

	// TODO 这里没有使用session去原生执行相应的sql语句,是否能够正常回滚有待测试
	err = l.svcCtx.Conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		selfContact := &user_contact.UserContact{}
		selfContact.OwnerId = util.ConvertInt64(req.OwnerId)
		selfContact.DstId = util.ConvertInt64(user.Id)
		_, selfErr := l.svcCtx.UserContactModel.Insert(l.ctx, selfContact)
		if selfErr != nil {
			return selfErr
		}

		targetContact := &user_contact.UserContact{}
		targetContact.OwnerId = util.ConvertInt64(user.Id)
		targetContact.DstId = util.ConvertInt64(req.OwnerId)
		_, targetErr := l.svcCtx.UserContactModel.Insert(l.ctx, targetContact)
		if targetErr != nil {
			return targetErr
		}

		return nil
	})
	if err != nil {
		return errors.New(exception.UserAddFailed.Code, exception.UserAddFailed.Msg)
	}

	return nil
}
