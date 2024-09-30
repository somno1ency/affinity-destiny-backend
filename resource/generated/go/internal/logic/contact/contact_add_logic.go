package contact

import (
	"context"

	"ad.com/resource/generated/go/internal/svc"
	"ad.com/resource/generated/go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return nil
}
