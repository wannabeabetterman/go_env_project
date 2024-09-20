package vo

type AirConditionerVo struct {
	Id                   int64  `json:"id,optional"`
	Name                 string `json:"name"`
	TenantCode           string `json:"tenantCode,optional"`
	StationCode          string `json:"stationCode,optional"`          // 站点code
	OverallState         string `json:"overallState,optional"`         // 整机状态0：停止;1 ：运行 2：故障
	ElectricHeatingState string `json:"electricHeatingState,optional"` // 电加热状态0：停止;1 ：运行 2：故障
	InternalFanState     string `json:"internalFanState,optional"`     // 内风机状态0：停止;1 ：运行 2：故障
	ExternalFanState     string `json:"externalFanState,optional"`     // 外风机状态0：停止;1 ：运行 2：故障
	CompressorState      string `json:"compressorState,optional"`      // 压缩机状态0：停止;1 ：运行 2：故障
	OutdoorTemperature   string `json:"outdoorTemperature,optional"`   // 室外温度
	IndoorTemperature    string `json:"indoorTemperature,optional"`    // 室内温度
	Humidity             string `json:"humidity,optional"`             // 湿度
}
