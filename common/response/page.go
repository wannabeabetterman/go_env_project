package response

type PageInfo struct {
	Page     int `form:"page,default=1"` //默认值
	PageSize int `form:"pageSize,default=10"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
