/*
*

	@author Yeoman
	@date:2022/7/14
	@Description
*/
package logic

import (
	svc "alger/common/context"
	"context"
)

// ApiLogic 实例化
type ApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ApiServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ApiServiceContext) ApiLogic {
	return ApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
