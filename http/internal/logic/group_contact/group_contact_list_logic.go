package group_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactListLogic {
	return &GroupContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactListLogic) GroupContactList(req *types.GroupIdReq) (resp []types.GroupContactResp, err error) {
	// todo: add your logic here and delete this line

	return
}