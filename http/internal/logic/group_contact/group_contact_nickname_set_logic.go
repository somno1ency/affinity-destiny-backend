package group_contact

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupContactNicknameSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupContactNicknameSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupContactNicknameSetLogic {
	return &GroupContactNicknameSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupContactNicknameSetLogic) GroupContactNicknameSet(req *types.GroupContactStringSetReq) error {
	// todo: add your logic here and delete this line

	return nil
}
