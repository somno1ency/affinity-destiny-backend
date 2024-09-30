package group

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	group_contact "ad.com/pkg/repo/group-contact"
	"ad.com/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type GroupJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupJoinLogic {
	return &GroupJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupJoinLogic) GroupJoin(req *types.GroupJoinReq) error {
	group, err := l.svcCtx.GroupModel.FindOneByGroupId(l.ctx, util.ConvertString(req.GroupId))
	if err != nil {
		return errors.New(exception.GroupNotFound.Code, exception.GroupNotFound.Msg)
	}
	userIds, err := l.svcCtx.GroupContactModel.FindUserIdsByGroupId(l.ctx, group.Id)
	if err != nil {
		return errors.New(exception.UnknownError.Code, err.Error())
	}
	if util.Contains(userIds, req.UserId) {
		return errors.New(exception.GroupAlreadyIn.Code, exception.GroupAlreadyIn.Msg)
	}

	groupContact := &group_contact.GroupContact{}
	groupContact.GroupId = util.ConvertInt64(group.Id)
	groupContact.UserId = util.ConvertInt64(req.UserId)
	_, err = l.svcCtx.GroupContactModel.Insert(l.ctx, groupContact)
	if err != nil {
		return errors.New(exception.GroupJoinFailed.Code, exception.GroupJoinFailed.Msg)
	}

	return nil
}
