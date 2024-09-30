package group

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	global "ad.com/pkg/const"
	"ad.com/pkg/exception"
	"ad.com/pkg/repo/group"
	group_contact "ad.com/pkg/repo/group-contact"
	"ad.com/pkg/util"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/x/errors"
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
	_, err := l.svcCtx.GroupModel.FindOneByName(l.ctx, util.ConvertString(req.Name))
	if err == nil {
		return errors.New(exception.GroupExists.Code, exception.GroupExists.Msg)
	}
	if err != sqlx.ErrNotFound {
		return errors.New(exception.UnknownError.Code, err.Error())
	}
	if count, err := l.svcCtx.GroupModel.CountByOwnerId(l.ctx, req.OwnerId); count >= global.GroupMaxQuantity {
		if err != nil {
			return errors.New(exception.GroupQualityLimited.Code, err.Error())
		}

		return errors.New(exception.GroupQualityLimited.Code, exception.GroupQualityLimited.Msg)
	}

	group := &group.Group{}
	_ = copier.Copy(group, req)
	group.GroupId = util.ConvertString(util.GenId())
	// TODO 这里没有使用session去原生执行相应的sql语句,是否能够正常回滚有待测试
	err = l.svcCtx.Conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		_, groupErr := l.svcCtx.GroupModel.Insert(l.ctx, group)
		if groupErr != nil {
			return groupErr
		}

		groupContact := &group_contact.GroupContact{}
		groupContact.GroupId = util.ConvertInt64(group.Id)
		groupContact.UserId = group.OwnerId
		_, groupContactErr := l.svcCtx.GroupContactModel.Insert(l.ctx, groupContact)
		if groupContactErr != nil {
			return groupContactErr
		}

		return nil
	})
	if err != nil {
		return errors.New(exception.GroupCreateFailed.Code, exception.GroupCreateFailed.Msg)
	}

	return nil
}
