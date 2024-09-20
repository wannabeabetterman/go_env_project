/*
*

	@author
	@date:
	@Description
*/
package handler

import (
	svc "alger/common/context"
	"alger/common/response"
	"alger/logic"
	login "alger/model/dto/login"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetLogin(svcCtx *svc.ApiServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req login.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.New("RequestParameterVerify")
			response.FailWithMessageId(w, err.Error())
			return
		}
		l := logic.NewApiLogic(r.Context(), svcCtx)
		data, err := l.GetLogin(req)
		if err != nil {
			response.FailWithMessageId(w, err.Error())
		} else {
			response.OkWithData(w, data)
		}
	}
}
