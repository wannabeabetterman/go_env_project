package logic

import (
	"alger/common/global"
	login "alger/model/dto/login"
	"alger/repository"
)

func (l *ApiLogic) GetLogin(req login.LoginReq) (reply login.LoginResp, err error) {

	db := global.Db
	resp, err := repository.GetUser(db, req)
	if resp.ID != 0 {
		return login.LoginResp{
			Result: "1",
		}, err
	}
	return login.LoginResp{
		Result: "0",
	}, err
}
