package user

import (
	"context"
	"fmt"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"
	"ad.com/pkg/shared"

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
	// code := util.GenId(6)
	code := "123456"
	logx.Infof("current code: %s", code)
	// TODO: send code to phone
	if err := l.svcCtx.RdsClient.Setex(fmt.Sprintf(shared.CodeSendToPhone, req.Mobile), code, shared.RdsCodeCacheTime); err != nil {
		logx.Errorf("set code to redis failed, err: %v", err)
		return &exception.CodeSendFailed
	}

	return nil
}
