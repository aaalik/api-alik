package autils

import (
	"encoding/json"
	"strconv"
)

func StringToFloat(val string, defValue float64) float64 {
	if result, err := strconv.ParseFloat(val, 64); err == nil {
		return result
	} else {
		return defValue
	}
}

func StringToInt(val string, defValue int) int {
	if result, err := strconv.Atoi(val); err == nil {
		return result
	} else {
		return defValue
	}
}

func StringToBoolean(val string, defValue bool) bool {
	if result, err := strconv.ParseBool(val); err == nil {
		return result
	} else {
		return defValue
	}
}

func FloatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 6, 64)
}

func IntToString(val int) string {
	return strconv.Itoa(val)
}

func StructToJson(p interface{}) string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}

	return string(b)
}