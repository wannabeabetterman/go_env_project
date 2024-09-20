package model

import "time"

// IecCemsAirConditioner 空调数据表
type AirConditioner struct {
	ID                         uint64    `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"id"`
	TenantCode                 string    `gorm:"column:tenant_code;type:varchar(32);not null" json:"tenantCode"`                                    // 租户编码
	StationCode                string    `gorm:"column:station_code;type:varchar(900);not null" json:"stationCode"`                                 // 站点code
	EquipmentCode              string    `gorm:"column:equipment_code;type:varchar(20);not null" json:"equipmentCode"`                              // 设备编码
	VersionInformation         string    `gorm:"column:version_information;type:varchar(20);not null" json:"versionInformation"`                    // 版本信息
	OverallState               string    `gorm:"column:overall_state;type:varchar(20);not null" json:"airConditionerSystemState"`                   // 整机状态0：停止;1 ：运行 2：故障
	InternalFanState           string    `gorm:"column:internal_fan_state;type:varchar(20);not null" json:"airConditionerFanState"`                 // 内风机状态0：停止;1 ：运行 2：故障
	ExternalFanState           string    `gorm:"column:external_fan_state;type:varchar(20);not null" json:"externalFanState"`                       // 外风机状态0：停止;1 ：运行 2：故障
	CompressorState            string    `gorm:"column:compressor_state;type:varchar(20);not null" json:"airConditionerCompressorState"`            // 压缩机状态0：停止;1 ：运行 2：故障
	ElectricHeatingState       string    `gorm:"column:electric_heating_state;type:varchar(20);not null" json:"airConditionerElectricHeatingState"` // 电加热状态0：停止;1 ：运行 2：故障
	EmergencyFanState          string    `gorm:"column:emergency_fan_state;type:varchar(20);not null" json:"emergencyFanState"`                     // 应急风机状态0：停止;1 ：运行 2：故障
	OutdoorTemperature         string    `gorm:"column:outdoor_temperature;type:varchar(20);not null" json:"airConditionerOutdoorTemperature"`      // 室外温度
	IndoorTemperature          string    `gorm:"column:indoor_temperature;type:varchar(20);not null" json:"airConditionerIndoorTemperature"`        // 室内温度
	Humidity                   string    `gorm:"column:humidity;type:varchar(20);not null" json:"airConditionerHumidity"`                           // 湿度
	FaultCode                  string    `gorm:"column:fault_code;type:varchar(20);not null" json:"airConditionerFaultCode"`                        // 空调故障码0：正常;;其他：故障
	EquipmentCoolingPoint      string    `gorm:"column:equipment_cooling_point;type:varchar(20);not null" json:"equipmentCoolingPoint"`             // 设备制冷点
	EquipmentCoolingHysteresis string    `gorm:"column:equipment_cooling_hysteresis;type:varchar(20);not null" json:"equipmentCoolingHysteresis"`   // 设备制冷回差
	EquipmentHeatingPoint      string    `gorm:"column:equipment_heating_point;type:varchar(20);not null" json:"equipmentHeatingPoint"`             // 设备加热点
	EquipmentHeatingHysteresis string    `gorm:"column:equipment_heating_hysteresis;type:varchar(20);not null" json:"equipmentHeatingHysteresis"`   // 设备加热回差
	CoilTemperature            string    `gorm:"column:coil_temperature;type:varchar(20);not null" json:"coilTemperature"`                          // 盘管温度
	CondensationTemperature    string    `gorm:"column:condensation_temperature;type:varchar(20);not null" json:"condensationTemperature"`          // 冷凝温度
	ExhaustTemperature         string    `gorm:"column:exhaust_temperature;type:varchar(20);not null" json:"exhaustTemperature"`                    // 排气温度
	Current                    string    `gorm:"column:current;type:varchar(20);not null" json:"current"`                                           // 电流
	AcVoltage                  string    `gorm:"column:ac_voltage;type:varchar(20);not null" json:"acVoltage"`                                      // 交流电压
	DcVoltage                  string    `gorm:"column:dc_voltage;type:varchar(20);not null" json:"dcVoltage"`                                      // 直流电压
	ReportTime                 time.Time `gorm:"column:report_time;type:datetime;not null" json:"reportTime"`                                       // 上报时间
}

// TableName get sql table name.获取数据库表名
func (m *AirConditioner) TableName() string {
	return "iec_cems_dw_air_conditioner"
}
