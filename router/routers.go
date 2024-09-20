/*
 * @Author: lihao lihao@ikbvip.com
 * @Date: 2023-05-04 14:38:44
 * @LastEditors: lihao lihao@ikbvip.com
 * @LastEditTime: 2024-03-21 14:45:02
 * @FilePath: \iec-em-alger-server\router\routers.go
 * @Description:
 */
package router

import (
	svc "alger/common/context"
	"alger/middleware/i18n"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ApiServiceContext) {
	routers := initWebRoutes(serverCtx)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{

				rest.ToMiddleware(func(next http.Handler) http.Handler {
					return i18n.I18n(next, serverCtx)
				}),
			},
			routers...,
		),
		rest.WithPrefix("/alger"),
	)

}

func initWebRoutes(serverCtx *svc.ApiServiceContext) (routes []rest.Route) {

	routes = append(routes, UserRouter(serverCtx)...) //辅助系统

	return routes
}
