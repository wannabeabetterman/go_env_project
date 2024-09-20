package router

import (
	svc "alger/common/context"
	"alger/handler"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func UserRouter(serverCtx *svc.ApiServiceContext) []rest.Route {
	return []rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/api/v1/user/login",
			Handler: handler.GetLogin(serverCtx),
		},
	}
}
