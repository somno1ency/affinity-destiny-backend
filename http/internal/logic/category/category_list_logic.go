package category

import (
	"context"
	"fmt"

	"ad.com/http/internal/svc"
	"ad.com/http/internal/types"
	"ad.com/pkg/exception"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList() (resp []types.CategoryResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("CategoryList")
	// l.svcCtx.RdsClient.SetnxEx("aaa", "a-value", 60)
	// l.svcCtx.RdsClient.SetnxExCtx(l.ctx, "aaac", "b-value", 60)
	// return errors.New(exception.CodeSendFailed.Code, exception.CodeSendFailed.Msg)
	// return nil, errors.New(exception.CodeSendFailed.Code, exception.CodeSendFailed.Msg)

	return nil, &exception.CodeSendFailed
}
