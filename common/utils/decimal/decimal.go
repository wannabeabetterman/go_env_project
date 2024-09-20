package decimal

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

func FloatDecimal(num float64, prec int) float64 {
	if num == 0 {
		return 0
	}
	value, _ := decimal.NewFromString(strconv.FormatFloat(num, 'f', 5, 64))
	a := value.Round(int32(prec))
	floatValue, _ := a.Float64()

	return floatValue
}

//func BuildEconomyCoal(value float64, prec int) float64 {
//	newDecimal, _ := decimal.NewFromString(strconv.FormatFloat(value*consts.EconomyCoal, 'f', 5, 64))
//	a := newDecimal.Round(int32(prec))
//	floatValue, _ := a.Float64()
//	return floatValue
//}
//
//func BuildTreePlanting(value float64, prec int) float64 {
//	newDecimal, _ := decimal.NewFromString(strconv.FormatFloat(value*consts.TreePlanting, 'f', 5, 64))
//	a := newDecimal.Round(int32(prec))
//	floatValue, _ := a.Float64()
//	return floatValue
//}
//
//func BuildCo2(value float64, prec int) float64 {
//	newDecimal, _ := decimal.NewFromString(strconv.FormatFloat(value*consts.Co2Value, 'f', 5, 64))
//	a := newDecimal.Round(int32(prec))
//	floatValue, _ := a.Float64()
//	return floatValue
//}
//
//func BuildSo2(value float64, prec int) float64 {
//	newDecimal, _ := decimal.NewFromString(strconv.FormatFloat(value*consts.So2Value, 'f', 5, 64))
//	a := newDecimal.Round(int32(prec))
//	floatValue, _ := a.Float64()
//	return floatValue
//}

func StringToKeepOneDecimalPlace(value string, prec int) string {
	if len(value) > 0 {
		num, _ := decimal.NewFromString(value)
		a := num.Round(int32(prec))
		floatValue, _ := a.Float64()
		targetFloatValue := strconv.FormatFloat(floatValue, 'f', prec, 64)
		return targetFloatValue
	}
	return ""
}

func F2i(f float64) int64 {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return int64(i)
}

// 保留两位小数
func UintToFeeFloat(u int64) float64 {

	floatValue := float64(u)

	a := FloatDecimal(floatValue/100, 2)
	return a
}

func StringToDecimalPlace(value string, prec int) float64 {
	newDecimal, _ := decimal.NewFromString(value)
	a := newDecimal.Round(int32(prec))
	floatValue, _ := a.Float64()
	return floatValue
}

func IntToDecimal(num int64, prec int) float64 {
	floatValue, err := strconv.ParseFloat(strconv.Itoa(int(num)), 64)
	if err != nil {
		return 0
	}
	return floatValue
}

func FloatToString(value float64, prec int) string {
	floatString := strconv.FormatFloat(value, 'f', prec, 64)
	newDecimal, _ := decimal.NewFromString(floatString)
	a := newDecimal.Round(int32(prec))
	floatValue, _ := a.Float64()
	return strconv.FormatFloat(floatValue, 'f', prec, 64)
}

func StringToKeepOneDecimal(value string, prec int) string {
	newDecimal, _ := decimal.NewFromString(value)
	a := newDecimal.Round(int32(prec))
	floatValue, _ := a.Float64()
	targetFloatValue := strconv.FormatFloat(floatValue, 'f', prec, 64)
	return targetFloatValue
}

func FloatToInt64(num float64) int64 {
	num = num * 100
	valueInt64 := F2i(num)
	return valueInt64
}

func IntToDecimalDivisor100(num int64, prec int) float64 {
	floatValue, err := strconv.ParseFloat(strconv.Itoa(int(num)), 64)
	if err != nil {
		return 0
	}
	newDecimal, _ := decimal.NewFromString(strconv.FormatFloat(floatValue/100, 'f', 5, 64))
	a := newDecimal.Round(int32(prec))
	floatValue, _ = a.Float64()
	return floatValue

}
