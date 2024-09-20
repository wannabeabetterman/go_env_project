package request

type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type ExportPageInfo struct {
	Page     int `json:"page,optional"`
	PageSize int `json:"pageSize,optional"`
}
