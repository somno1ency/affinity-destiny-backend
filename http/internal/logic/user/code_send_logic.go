package user

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CodeSendLogic {
	return &CodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CodeSendLogic) CodeSend(req *types.CodeSendReq) error {
	// todo: add your logic here and delete this line

	return nil
}
