package group

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type GroupLoadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupLoadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupLoadLogic {
	return &GroupLoadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupLoadLogic) GroupLoad(req *types.PathIdReq) (resp []types.GroupResp, err error) {
	groupIds, err := l.svcCtx.GroupContactModel.FindGroupIdsByUserId(l.ctx, req.Id)
	if err != nil {
		return nil, errors.New(exception.UnknownError.Code, err.Error())
	}

	resp = make([]types.GroupResp, 0)
	if len(groupIds) == 0 {
		return resp, nil
	}
	groups, _ := l.svcCtx.GroupModel.FindInIds(l.ctx, groupIds)
	_ = copier.Copy(&resp, groups)

	return resp, nil
}
