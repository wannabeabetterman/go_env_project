package dto

type AirConditionerQueryReq struct {
	Id                int64  `json:"id,optional"`
	StationCode       string `json:"stationCode,optional"`
	TenantCode        string `json:"tenantCode,optional"` //租户编码
	EquipmentCode     string `json:"equipmentCode,optional"`
	EquipmentTypeCode string `json:"equipmentTypeCode,optional"` //租户编码
}

type TemperatureAndHumidityStatisticReply struct {
	OutdoorTemperature string `gorm:"column:outdoorTemperature" json:"outdoorTemperature"` // 室外温度
	IndoorTemperature  string `gorm:"column:indoorTemperature" json:"indoorTemperature"`   // 室内温度
	Humidity           string `gorm:"column:humidity" json:"humidity"`                     // 湿度
	Date               string `gorm:"column:date" json:"date"`
}
