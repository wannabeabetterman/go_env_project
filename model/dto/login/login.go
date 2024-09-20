package dto

type LoginReq struct {
	UserName string `json:"userName,optional"`

	Password string `json:"password,optional"`
}
type LoginResp struct {
	Result string `json:"result,optional"`
}
