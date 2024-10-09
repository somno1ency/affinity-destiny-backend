package user

import (
	"context"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CodeValidateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCodeValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CodeValidateLogic {
	return &CodeValidateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CodeValidateLogic) CodeValidate(req *types.CodeValidateReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
